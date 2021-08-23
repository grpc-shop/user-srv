package dao

import (
	"github.com/grpc-shop/user-srv/model"
	"gorm.io/gorm"
)

type UserDao interface {
	CreateUser(user model.User) (uId int64, err error)
}

var _ UserDao = (*UserImpl)(nil)

type UserImpl struct {
	db *gorm.DB
}

func NewUserImpl(db *gorm.DB) UserDao {
	return &UserImpl{
		db: db,
	}
}

func (u UserImpl) CreateUser(user model.User) (uId int64, err error) {
	base := u.db.Create(&user)
	err = base.Error
	uId = user.Id
	return
}
