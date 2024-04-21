package entities

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name          string `gorm:"type:varchar(50)"`
	EstablishDate *time.Time
	StudyField    string `gorm:"type:varchar(50)"`
	ShortHandName string `gorm:"type:varchar(5)"`
	ImageUrl      string `gorm:"type:varchar(1000)"`
}
