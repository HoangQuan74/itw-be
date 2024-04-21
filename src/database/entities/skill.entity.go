package entities

import (
	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name string `gorm:"type:varchar(100)"`
}
