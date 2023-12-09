package repository

import (
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	First(cond ...any) (user models.User, err error)
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		DB: db,
	}
}

func (uR UserRepo) First(cond ...any) (user models.User, err error) {
	err = uR.DB.First(&user, cond...).Error
	return
}
