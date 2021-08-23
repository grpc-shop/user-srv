package service

import (
	"github.com/grpc-shop/user-srv/dao"
	"github.com/grpc-shop/user-srv/model"
)

type UserServer interface {
	CreateUser(user model.User) (uId int64, err error)
}

var _ UserServer = (*UserServerImpl)(nil)

type UserServerImpl struct {
	dao dao.UserDao
}

func NewUserServer(dao dao.UserDao) UserServer {
	return &UserServerImpl{dao: dao}
}

func (u *UserServerImpl) CreateUser(user model.User) (uId int64, err error) {
	return u.dao.CreateUser(user)
}
