package model

type Gender struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(16);unique" json:"name"`
}

type Genders []Gender

func (Gender) TableName() string {
	return "master.gender"
}
