package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	StudentID  uuid.UUID `gorm:"size:255"`
	BusinessID uuid.UUID `gorm:"size:255"`

	Student  Student
	Business Business
}
