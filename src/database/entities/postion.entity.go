package entities

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}
