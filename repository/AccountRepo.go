package repository

import (
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepo struct {
	DB *gorm.DB
}

type AccountRepoInterface interface {
	Find(cond ...any) (rows []models.Account, err error)
	First(cond ...any) (account models.Account, err error)
	Create(account *models.Account) (err error)
}

func NewAccountRepo(db *gorm.DB) AccountRepoInterface {
	return &AccountRepo{
		DB: db,
	}
}

func (uR AccountRepo) Find(cond ...any) (rows []models.Account, err error) {
	err = uR.DB.Find(&rows, cond...).Error
	return
}

func (uR AccountRepo) First(cond ...any) (account models.Account, err error) {
	err = uR.DB.Preload(clause.Associations).First(&account, cond...).Error
	return
}

func (uR AccountRepo) Create(account *models.Account) (err error) {
	err = uR.DB.Clauses(clause.Returning{}).Create(&account).Error
	return
}
