//+build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/grpc-shop/user-srv/dao"
	"github.com/grpc-shop/user-srv/service"
	"gorm.io/gorm"
)

func InitUserHandler(db *gorm.DB) *UserHandler {
	panic(wire.Build(
		dao.NewUserImpl,
		service.NewUserServer,
		NewUserHandler,
	))
	return &UserHandler{}
}
