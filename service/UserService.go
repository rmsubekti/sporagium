package service

import (
	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/repository"
	"github.com/rmsubekti/sporagium/utils/srx"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepoInterface
}
type UserServiceInterface interface {
	Create(user *models.User) (err error)
	FirstById(id string) (user models.User, err error)
	First(cond ...any) (user models.User, err error)
	Find(cond ...any) (users []models.User, err error)
	Login(login dto.Login) (user models.User, err error)
	Register(user *models.User) (err error)
}

func NewUserService() UserServiceInterface {
	return &userService{
		repo: repository.NewUserRepo(db),
	}
}

func (u userService) Create(user *models.User) (err error) {
	return u.repo.Create(user)
}

func (u userService) FirstById(id string) (user models.User, err error) {
	return u.repo.First("id = ?", id)
}

func (u userService) First(cond ...any) (user models.User, err error) {
	return u.repo.First(cond...)
}

func (u userService) Find(cond ...any) (users []models.User, err error) {
	return u.repo.Find(cond...)
}

func (u userService) Login(login dto.Login) (user models.User, err error) {
	if srx.Email(login.Email).Ok() {
		if user, err = u.repo.First("email = ?", login.Email); err != nil {
			return
		}
	} else if user, err = u.repo.First("user_name = ?", login.Email); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return
	}
	return
}
func (u userService) Register(user *models.User) (err error) {
	if err = user.Validate(); err != nil {
		return
	}
	if err = user.GenerateHashPassword(); err != nil {
		return
	}
	user.ID = uuid.New()
	if err = u.repo.Create(user); err != nil {
		return
	}
	return
}
