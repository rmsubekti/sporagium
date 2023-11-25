package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/utils/srx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `gorm:"type:varchar(120)" json:"name"`
	Email     string    `gorm:"type:varchar(180);unique" json:"email"`
	Phone     string    `gorm:"type:varchar(15);unique" json:"phone"`
	UserName  string    `gorm:"type:varchar(80);unique" json:"user_name"`
	Password  string    `gorm:"type:text;unique" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Users []User

func (User) TableName() string {
	return "user.user"
}

func (u User) Validate() (err error) {
	if err = srx.Email(u.Email).Validate(); err != nil {
		return
	}
	if err = srx.Password(u.Password).Validate(); err != nil {
		return
	}
	if err = srx.Phone(u.Phone).Validate(); err != nil && len(u.Phone) > 0 {
		return
	}
	return
}

func (u *User) GenerateHashPassword() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
	return
}
