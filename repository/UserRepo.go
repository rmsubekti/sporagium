package repository

import (
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	Find(cond ...any) (rows []models.User, err error)
	First(cond ...any) (user models.User, err error)
	Create(user *models.User) (err error)
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &UserRepo{
		DB: db,
	}
}

func (uR UserRepo) Find(cond ...any) (rows []models.User, err error) {
	err = uR.DB.Find(&rows, cond...).Error
	return
}

func (uR UserRepo) First(cond ...any) (user models.User, err error) {
	err = uR.DB.First(&user, cond...).Error
	return
}

func (uR UserRepo) Create(user *models.User) (err error) {
	err = uR.DB.Clauses(clause.Returning{}).Create(&user).Error
	return
}
