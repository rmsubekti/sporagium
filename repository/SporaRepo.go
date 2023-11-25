package repository

import (
	"github.com/rmsubekti/sporagium/helper"
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SporaRepo struct {
	DB *gorm.DB
}

type SporaRepoInterface interface {
	Find(cond ...any) (rows []models.Spora, err error)
	First(cond ...any) (Spora models.Spora, err error)
	Create(Spora *models.Spora) (err error)
	Delete(Spora *models.Spora) (err error)
	Update(Spora *models.Spora) (err error)
	Paginate(paginator *helper.Paginator, userID string) (err error)
}

func NewSporaRepo(db *gorm.DB) SporaRepoInterface {
	return &SporaRepo{
		DB: db,
	}
}

func (sR SporaRepo) Find(cond ...any) (rows []models.Spora, err error) {
	err = sR.DB.Find(&rows, cond...).Error
	return
}

func (sR SporaRepo) First(cond ...any) (Spora models.Spora, err error) {
	err = sR.DB.Preload(clause.Associations).First(&Spora, cond...).Error
	return
}

func (sR SporaRepo) Create(spora *models.Spora) (err error) {
	err = sR.DB.Clauses(clause.Returning{}).Create(&spora).Error
	return
}

func (sR SporaRepo) Delete(spora *models.Spora) (err error) {
	err = sR.DB.Delete(&spora).Error
	return
}
func (sR SporaRepo) Update(spora *models.Spora) (err error) {
	temp := &models.Spora{}
	*temp = *spora
	err = sR.DB.Omit(clause.Associations).Model(&spora).Clauses(clause.Returning{}).Updates(&temp).Error
	return
}

func (sR SporaRepo) Paginate(paginator *helper.Paginator, userID string) (err error) {
	var data []models.Spora
	countDB := sR.DB.Group("id").Model(&models.Spora{})
	rowsDB := sR.DB.Scopes(paginator.Scopes()).Where("user_id = ?", userID).Preload(clause.Associations)

	if err = paginator.SetCount(countDB); err != nil {
		return
	}
	if err = rowsDB.Find(&data).Error; err != nil {
		return
	}
	paginator.Paginate(data)
	return
}
