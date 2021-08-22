package service

import (
	"github.com/grpc-shop/user-srv/dao"
)

type UserServer interface {
}

var _ UserServer = (*UserServerImpl)(nil)

type UserServerImpl struct {
	dao dao.UserDao
}

func NewUserServer(dao dao.UserDao) UserServer {
	return &UserServerImpl{dao: dao}
}
