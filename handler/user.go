package handler

import (
	"context"
	"github.com/grpc-shop/user-srv/proto/user"
	"github.com/grpc-shop/user-srv/service"
)

var _ user.UserServer = (*UserHandler)(nil)

type UserHandler struct {
	user.UnimplementedUserServer
	server service.UserServer
}

func NewUserHandler(server service.UserServer) *UserHandler {
	return &UserHandler{server: server}
}

func (u UserHandler) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserReply, error) {
	panic("implement me")
}
