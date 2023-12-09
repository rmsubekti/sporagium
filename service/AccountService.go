package service

import (
	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/repository"
	"github.com/rmsubekti/sporagium/utils/srx"
	"golang.org/x/crypto/bcrypt"
)

type accountService struct {
	repo repository.AccountRepoInterface
}
type AccountServiceInterface interface {
	Create(account *models.Account) (err error)
	FirstById(id string) (account models.Account, err error)
	First(cond ...any) (account models.Account, err error)
	Find(cond ...any) (accounts []models.Account, err error)
	Login(login dto.Login) (account models.Account, err error)
	Register(account *models.Account) (err error)
}

func NewAccountService() AccountServiceInterface {
	return &accountService{
		repo: repository.NewAccountRepo(db),
	}
}

func (a accountService) Create(account *models.Account) (err error) {
	return a.repo.Create(account)
}

func (a accountService) FirstById(id string) (account models.Account, err error) {
	return a.repo.First("id = ?", id)
}

func (a accountService) First(cond ...any) (account models.Account, err error) {
	return a.repo.First(cond...)
}

func (a accountService) Find(cond ...any) (accounts []models.Account, err error) {
	return a.repo.Find(cond...)
}

func (a accountService) Login(login dto.Login) (account models.Account, err error) {
	if srx.Email(login.Email).Ok() {
		if account, err = a.repo.First("email = ?", login.Email); err != nil {
			return
		}
	} else if account, err = a.repo.First("user_name = ?", login.Email); err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(login.Password)); err != nil {
		return
	}
	return
}
func (a accountService) Register(account *models.Account) (err error) {
	if err = account.Validate(); err != nil {
		return
	}
	if err = account.GenerateHashPassword(); err != nil {
		return
	}
	account.ID = uuid.New()
	account.User.ID = account.ID
	if err = a.repo.Create(account); err != nil {
		return
	}
	return
}
