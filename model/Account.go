package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rmsubekti/sporagium/dto"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string    `gorm:"type:varchar(120);unique"`
	Password     string    `gorm:"type:varchar(200)" json:"password"`
	RegisteredAt time.Time `gorm:"autoCreateTime" json:"registered_at"`
	User         User      `gorm:"foreignKey:ID;references:ID"`
	Emails       Emails    `gorm:"foreignKey:UserID;references:ID"`
	Phones       Phones    `gorm:"foreignKey:UserID;references:ID"`
}
type Accounts []Account

func (Account) TableName() string {
	return "user.account"
}
func (a *Account) Set(r dto.Register) *Account {
	id := uuid.New()
	primary := true
	a.ID = id
	a.Password = r.Password
	a.Username = r.Username
	a.Emails = Emails{
		{UserID: id, Address: r.Email, Primary: &primary},
	}
	a.Phones = Phones{
		{UserID: id, Number: r.Phone, Primary: &primary},
	}
	a.User.ID = id
	a.User.GenderID = 1
	a.User.Name = r.Name
	return a
}

func (a *Account) Create() (err error) {
	if err = a.Emails.Validate(); err != nil {
		return
	}
	if err = a.Phones.Validate(); err != nil {
		return
	}
	if err = a.Emails.Registered(); err != nil {
		return
	}
	if err = a.Phones.Registered(); err != nil {
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	a.Password = string(hashedPassword)

	if err = db.Create(&a).Error; err != nil {
		return
	}
	return
}

func (a *Account) Login(cred string, password string) error {
	var mail Email
	if len(cred) < 1 {
		return errors.New("empty email or username")
	}

	if rxMail.MatchString(cred) {
		if err := mail.Get(cred); err != nil {
			return errors.New("email is not registered")
		}
		if err := a.Find(mail.UserID); err != nil {
			return err
		}
	} else {
		if err := a.Get(cred); err != nil {
			return err
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password)); err != nil {
		return errors.New("wrong password")
	}
	return nil

}

func (a *Account) Find(id uuid.UUID) error {
	return db.Preload("User").First(a, "id=?", id).Error
}

func (a *Account) Get(username string) error {
	return db.Preload("User").First(a, "username=?", username).Error
}
