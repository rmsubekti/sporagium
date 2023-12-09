package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/utils/srx"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Email     string    `gorm:"type:varchar(180);unique" json:"email"`
	Phone     string    `gorm:"type:varchar(15);unique" json:"phone"`
	UserName  string    `gorm:"type:varchar(80);unique" json:"user_name"`
	User      User      `gorm:"foreignKey:ID;references:ID" json:"user"`
	Password  string    `gorm:"type:text;unique" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Accounts []Account

func (Account) TableName() string {
	return "user.account"
}

func (a Account) Validate() (err error) {
	if err = srx.Email(a.Email).Validate(); err != nil {
		return
	}
	if err = srx.Password(a.Password).Validate(); err != nil {
		return
	}
	if err = srx.Phone(a.Phone).Validate(); err != nil && len(a.Phone) > 0 {
		return
	}
	return
}

func (a *Account) GenerateHashPassword() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	a.Password = string(hash)
	return
}
