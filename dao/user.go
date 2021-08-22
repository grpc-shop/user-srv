package dao

import (
	"gorm.io/gorm"
)

type UserDao interface {
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
