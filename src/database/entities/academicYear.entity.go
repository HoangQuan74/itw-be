package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AcademicYear struct {
	gorm.Model
	CurrentYear int
	SchoolID    uuid.UUID `gorm:"size:255"`

	School School
}
