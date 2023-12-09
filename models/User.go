package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	UserName string    `gorm:"type:varchar(80)" json:"user_name"`
	Name     string    `gorm:"type:varchar(100)" json:"name"`
}

func (User) TableName() string {
	return "user.user"
}
