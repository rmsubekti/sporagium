package service

import (
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/repository"
)

type userService struct {
	repo repository.UserRepoInterface
}
type UserServiceInterface interface {
	FirstById(id string) (user models.User, err error)
}

func NewUserService() UserServiceInterface {
	return &userService{
		repo: repository.NewUserRepo(db),
	}
}

func (a userService) FirstById(id string) (user models.User, err error) {
	return a.repo.First("id = ?", id)
}
