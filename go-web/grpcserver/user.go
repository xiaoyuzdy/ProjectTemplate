package grpcserver

import (
	"context"
	"go-web/models"
	"go-web/proto"
)

type HelloService struct {
}

func (h *HelloService) Hello(ctx context.Context, args *proto.String) (*proto.String, error) {
	return &proto.String{
		Value: "receive: " + args.Value + "  send: hi,client",
	}, nil
}

func (h *HelloService) GetUserInfo(ctx context.Context, args *proto.String) (*proto.User, error) {
	u := models.User{}
	err := models.UserHandler.QueryLastByWhere(&u, "account = ?", args.Value)
	if err != nil {
		return nil, err
	}
	return &proto.User{
		UserName:    u.UserName,
		AccountType: int32(u.AccountType),
	}, nil
}
