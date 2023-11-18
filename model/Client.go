package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

type Client struct {
	ID      string         `gorm:"type:text;primaryKey" json:"id"`
	SporaID uuid.UUID      `json:"spora_id"`
	Secret  string         `gorm:"type:text" json:"secret"`
	Domain  string         `gorm:"type:text" json:"domain"`
	Data    datatypes.JSON `json:"data"`
}

type Clients []Client

func (Client) TableName() string {
	return "spora.client"
}

func (c *Client) Create() (err error) {
	c.ID = uuid.NewString()
	c.Secret = uuid.NewString()
	if err = db.Clauses(clause.Returning{}).Create(&c).Error; err != nil {
		return
	}
	return
}
