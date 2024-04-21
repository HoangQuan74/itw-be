package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20)"`
	SchoolID uuid.UUID `gorm:"type:uuid"`

	School School
}
