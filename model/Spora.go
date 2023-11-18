package model

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

var rxUrl = regexp.MustCompile(`(?m)(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z]{2,}(\.[a-zA-Z]{2,})(\.[a-zA-Z]{2,})?\/[a-zA-Z0-9]{2,}|((https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z]{2,}(\.[a-zA-Z]{2,})(\.[a-zA-Z]{2,})?)|(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)?[a-zA-Z0-9]{2,}\.[a-zA-Z0-9]{2,}\.[a-zA-Z0-9]{2,}(\.[a-zA-Z0-9]{2,})? `)

type Spora struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	UserID        uuid.UUID `json:"-"`
	ClientSecrets Clients   `gorm:"foreignKey:SporaID" json:"clients"`
	Name          string    `gorm:"type:varchar(125)" json:"name"`
	HomePage      string    `gorm:"type:varchar(200);unique" json:"homepage"`
	Description   string    `gorm:"type:varchar(255)" json:"description"`
	CallbackURL   string    `gorm:"type:varchar(200);unique" json:"callback"`
}
type Sporas []Spora

func (Spora) TableName() string {
	return "spora.spora"
}

func (s Spora) Validate() error {
	if !strings.HasPrefix(s.CallbackURL, "http") {
		s.CallbackURL = "https://" + s.CallbackURL
	}

	if !rxUrl.MatchString(s.CallbackURL) {
		return errors.New("callback url is not valid url")
	}
	return nil
}

func (s *Spora) Create(userID string) (err error) {
	if err = s.Validate(); err != nil {
		return
	}
	s.ID = uuid.New()
	s.UserID = uuid.MustParse(userID)
	return db.Clauses(clause.Returning{}).Create(&s).Error
}

func (s *Spora) Update(id string) (err error) {
	temp := &Spora{}
	*temp = *s
	s.ID = uuid.MustParse(id)
	return db.Model(&s).Clauses(clause.Returning{}).Updates(&temp).Error
}

func (s *Spora) Get(id string) (err error) {
	return db.First(&s, "id = ?", id).Error
}

func (s *Sporas) GetAll(userId string) error {
	return db.Preload(clause.Associations).Find(s, "user_id=?", userId).Error
}
