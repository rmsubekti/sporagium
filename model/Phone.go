package model

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

var rxPhone = regexp.MustCompile(`^(?:[[:digit:]]+)$`)

type Phone struct {
	ID      uint      `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Number  string    `gorm:"type:varchar(18);unique" json:"phone"`
	Primary *bool     `json:"primary"`
}
type Phones []Phone

func (Phone) TableName() string {
	return "user.phone"
}
func (p *Phone) Validate() error {
	if p.Number != "" {
		if len(p.Number) < 10 {
			return errors.New("doest not correct phone number")
		}
		if !rxPhone.MatchString(p.Number) {
			return errors.New("number must be filled with only numbers")
		}
	}
	return nil
}

func (p *Phones) Validate() (err error) {
	for _, v := range *p {
		if err = v.Validate(); err != nil {
			break
		}
	}
	return
}
func (p Phone) Registered() error {
	if p.Number != "" && db.First(&Phone{}, "number =? ", p.Number).RowsAffected > 0 {
		return errors.New("phone already registered")
	}
	return nil
}

func (p *Phones) Registered() (err error) {
	for _, v := range *p {
		if err = v.Registered(); err != nil {
			break
		}
	}
	return
}
func (p *Phone) Create() (err error) {
	if err = db.Clauses(clause.Returning{}).Create(&p).Error; err != nil {
		return
	}
	return
}

func (p *Phone) Get(id uint) error {
	return db.First(&p, "id=?", id).Error
}
func (p *Phone) GetPrimary(userID string) error {
	return db.First(&p, "primary=true and user_id=?", userID).Error
}

func (p *Phones) List(userID string) error {
	return db.Find(&p, "user_id=?", userID).Error
}

func (p *Phone) Delete(id uint) error {
	return db.Delete(p, "id=?", id).Error
}
