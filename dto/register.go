package dto

import "github.com/rmsubekti/sporagium/models"

type Register struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password,omitempty"`
}

func (r Register) GetAccount() models.Account {
	return models.Account{
		UserName: r.UserName,
		Email:    r.Email,
		Phone:    r.Phone,
		Password: r.Password,
		User: models.User{
			UserName: r.UserName,
			Name:     r.Name,
		},
	}
}
