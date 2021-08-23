package handler

import (
	"context"
	"errors"
	"github.com/grpc-shop/user-srv/model"
	"github.com/grpc-shop/user-srv/proto/user"
	"github.com/grpc-shop/user-srv/service"
)

var _ user.UserServer = (*UserHandler)(nil)

var (
	PasswordErr = errors.New("两次密码不一致")
)

type UserHandler struct {
	user.UnimplementedUserServer
	server service.UserServer
}

func NewUserHandler(server service.UserServer) *UserHandler {
	return &UserHandler{server: server}
}

func (u UserHandler) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserReply, error) {
	resp := user.CreateUserReply{
		Code: user.Code_Success,
		Data: new(user.CreateUserReplyUser),
	}
	if req.GetPassword() != req.GetPasswordAgain() {
		resp.Code = user.Code_CreateErr
		return &resp, PasswordErr
	}
	var (
		userModel model.User
	)
	userModel.Name = req.GetName()
	userModel.Email = req.GetEmail()
	userModel.Password = req.GetPassword()
	id, err := u.server.CreateUser(userModel)
	if err != nil {
		resp.Code = user.Code_CreateErr
		return &resp, err
	}
	resp.Data.UserId = id
	return &resp, nil
}
