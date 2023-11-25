package service

import (
	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/repository"
)

type secretService struct {
	repo repository.SecretRepoInterface
}

type SecretServiceInterface interface {
	FirstByID(id string) (secret models.Secret, err error)
	FindByUserID(userID string) (secret []models.Secret, err error)
	Generate(sporaID uuid.UUID) (err error)
	Delete(ID uint) (err error)
}

func NewSecretService() SecretServiceInterface {
	return &secretService{
		repo: repository.NewSecretRepo(db),
	}
}
func (s secretService) FirstByID(id string) (secret models.Secret, err error) {
	return s.repo.First("id = ?", id)
}
func (s secretService) FindByUserID(userID string) (secret []models.Secret, err error) {
	return s.repo.Find("user_id = ?", userID)
}
func (s secretService) Generate(sporaID uuid.UUID) (err error) {
	secret := models.Secret{
		SporaID: sporaID,
		Secret:  uuid.NewString(),
	}
	return s.repo.Create(&secret)
}

func (s secretService) Delete(ID uint) (err error) {
	return s.repo.Delete(&models.Secret{ID: ID})
}
