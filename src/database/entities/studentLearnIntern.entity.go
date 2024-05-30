package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentLearnIntern struct {
	gorm.Model
	StudentID    uuid.UUID `gorm:"size:255"`
	Score        int       `gorm:"type:tinyint"`
	BoardID      uuid.UUID `gorm:"size:255"`
	SubjectID    uuid.UUID `gorm:"size:255"`
	PassedStatus string
	RegistStatus string

	Subject InternSubject
}
