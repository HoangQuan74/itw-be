package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100)"`
	SchoolID uuid.UUID `gorm:"type:uuid"`

	School School
}
