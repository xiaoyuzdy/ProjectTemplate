package cmd

import (
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"go-web/component/log"
	"go-web/grpcserver"
	"go-web/proto"
	"google.golang.org/grpc"
	"net"
)

func Grpc(ctx *cli.Context) {
	grpcServer := grpc.NewServer()
	proto.RegisterHelloServiceServer(grpcServer, new(grpcserver.HelloService))
	address := ":" + viper.GetString("system.rpcport")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Sugar.Fatal(err)
	}
	grpcServer.Serve(lis)
}
