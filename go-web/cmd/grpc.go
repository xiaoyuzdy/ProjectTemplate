package cmd

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"go-web/component/log"
	"go-web/grpcserver"
	"go-web/proto"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func Grpc(ctx *cli.Context) {
	go startGrpc()
	err := startGateway
	if err != nil {
		log.Sugar.Error(err())
	}
}

func startGrpc() {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, new(grpcserver.HelloService))
	address := ":" + viper.GetString("system.rpcport")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Sugar.Fatal(err)
	}
	grpcServer.Serve(lis)
}

func startGateway() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	address := ":" + viper.GetString("system.rpcport")
	err := proto.RegisterHelloServiceHandlerFromEndpoint(
		ctx, mux, "127.0.0.1"+address,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return err
	}
	//监听的端口不能echo所用端口一致
	return http.ListenAndServe(":8090", mux)
}
