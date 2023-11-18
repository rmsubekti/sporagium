package model

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

var rxMail = regexp.MustCompile(`^(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+@(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+\.(?:[[:alpha:]]{2,6})$`)

type Email struct {
	ID      uint      `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Address string    `gorm:"type:varchar(120);unique" json:"email"`
	Primary *bool     `json:"primary"`
}
type Emails []Email

func (Email) TableName() string {
	return "user.email"
}
func (e *Email) Validate() error {
	if !rxMail.MatchString(e.Address) {
		return errors.New("invalid email address")
	}
	return nil
}

func (e *Emails) Validate() (err error) {
	for _, v := range *e {
		if err = v.Validate(); err != nil {
			break
		}
	}
	return
}

func (e *Email) Registered() error {
	if e.Address != "" && db.First(&Email{}, "address =? ", e.Address).RowsAffected > 0 {
		return errors.New("email address already registered")
	}
	return nil
}

func (e *Emails) Registered() (err error) {
	for _, v := range *e {
		if err = v.Registered(); err != nil {
			break
		}
	}
	return
}
func (e *Email) Create() (err error) {
	if err = db.Clauses(clause.Returning{}).Create(&e).Error; err != nil {
		return
	}
	return
}
func (e *Email) Get(email string) error {
	return db.First(&e, "address=?", email).Error
}

func (e *Email) GetPrimary(userID uuid.UUID) error {
	return db.First(&e, "primary=true and user_id=?", userID).Error
}

func (e *Emails) List(userID uuid.UUID) error {
	return db.Find(&e, "user_id=?", userID).Error
}

func (e *Email) Delete(id uint) error {
	return db.Delete(e, "id=?", id).Error
}
