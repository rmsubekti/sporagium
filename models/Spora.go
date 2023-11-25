package models

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Spora struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	User        User      `gorm:"foreignKey:UserID" json:"user"`
	UserID      uuid.UUID `json:"-"`
	Secrets     []Secret  `gorm:"foreignKey:SporaID" json:"secrets,omitempty"`
	Name        string    `gorm:"type:varchar(125)" json:"name"`
	HomePage    string    `gorm:"type:varchar(200);unique" json:"homepage"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CallbackURL string    `gorm:"type:text;unique" json:"callback_url"`
	Data        SporaData `gorm:"type:jsonb" json:"data"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
type Sporas []Spora

func (Spora) TableName() string {
	return "spora.spora"
}

func (s *Spora) Validate() (err error) {
	if !strings.HasPrefix(s.CallbackURL, "http") {
		return errors.New("callback url should start with http or https")
	}
	if len(s.Name) < 1 {
		return errors.New("name is required")
	}
	if len(s.HomePage) < 1 {
		return errors.New("app homepage is required")
	}
	return
}
