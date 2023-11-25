package service

import (
	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/helper"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/repository"
)

type sporaService struct {
	repo repository.SporaRepoInterface
}

type SporaServiceInteface interface {
	FirstByID(id string) (spora models.Spora, err error)
	FindByUserID(userID string) (spora []models.Spora, err error)
	Create(spora *models.Spora, userID string) (err error)
	Update(spora *models.Spora, ID string) (err error)
	Paginate(paginator *helper.Paginator, userID string) (err error)
}

func NewSporaService() SporaServiceInteface {
	return &sporaService{
		repo: repository.NewSporaRepo(db),
	}
}

func (s sporaService) FirstByID(id string) (spora models.Spora, err error) {
	return s.repo.First("id = ?", id)
}
func (s sporaService) FindByUserID(userID string) (spora []models.Spora, err error) {
	return s.repo.Find("user_id = ?", userID)
}
func (s sporaService) Create(spora *models.Spora, userID string) (err error) {
	if err = spora.Validate(); err != nil {
		return
	}
	spora.UserID = uuid.MustParse(userID)
	return s.repo.Create(spora)
}
func (s sporaService) Update(spora *models.Spora, ID string) (err error) {
	if err = spora.Validate(); err != nil {
		return
	}
	spora.ID = uuid.MustParse(ID)
	return s.repo.Update(spora)
}
func (s sporaService) Paginate(paginator *helper.Paginator, userID string) (err error) {
	return s.repo.Paginate(paginator, userID)
}
