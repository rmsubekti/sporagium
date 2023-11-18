package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name     string    `gorm:"type:varchar(150)" json:"name"`
	Gender   Gender    `gorm:"foreignKey:GenderID" json:"gender"`
	GenderID uint      `json:"gender_id,omitempty"`
}

type Users []User

func (User) TableName() string {
	return "user.user"
}

func (u *User) Get(id string) (err error) {
	if err = db.Preload(clause.Associations).First(&u, "id=?", id).Error; err != nil {
		return
	}
	return
}
