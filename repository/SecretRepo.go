package repository

import (
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SecretRepo struct {
	DB *gorm.DB
}

type SecretRepoInterface interface {
	Find(cond ...any) (rows []models.Secret, err error)
	First(cond ...any) (Secret models.Secret, err error)
	Create(secret *models.Secret) (err error)
	Delete(secret *models.Secret) (err error)
}

func NewSecretRepo(db *gorm.DB) SecretRepoInterface {
	return &SecretRepo{
		DB: db,
	}
}

func (sR SecretRepo) Find(cond ...any) (rows []models.Secret, err error) {
	err = sR.DB.Find(&rows, cond...).Error
	return
}

func (sR SecretRepo) First(cond ...any) (secret models.Secret, err error) {
	err = sR.DB.Preload(clause.Associations).First(&secret, cond...).Error
	return
}

func (sR SecretRepo) Create(spora *models.Secret) (err error) {
	err = sR.DB.Clauses(clause.Returning{}).Create(&spora).Error
	return
}

func (sR SecretRepo) Delete(spora *models.Secret) (err error) {
	err = sR.DB.Delete(&spora).Error
	return
}
