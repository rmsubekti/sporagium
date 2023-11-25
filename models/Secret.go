package models

import (
	"github.com/google/uuid"
)

type Secret struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	SporaID uuid.UUID `json:"spora_id"`
	Secret  string    `gorm:"type:text" json:"secret"`
}

type Clients []Secret

func (Secret) TableName() string {
	return "spora.secret"
}

// func (c *Secret) Create() (err error) {
// 	c.ID = uuid.NewString()
// 	c.Secret = uuid.NewString()
// 	if err = db.Clauses(clause.Returning{}).Create(&c).Error; err != nil {
// 		return
// 	}
// 	return
// }
