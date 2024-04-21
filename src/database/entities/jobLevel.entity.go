package entities

import (
	"gorm.io/gorm"
)

type JobLevel struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}
