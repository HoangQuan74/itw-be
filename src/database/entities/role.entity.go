package entities

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20);unique"`
	Description string `gorm:"type:varchar(100)"`
}
