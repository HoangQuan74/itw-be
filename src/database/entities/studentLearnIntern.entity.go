package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentLearnIntern struct {
	gorm.Model
	StudentID    uuid.UUID
	Score        int `gorm:"type:tinyint"`
	BoardID      uuid.UUID
	SubjectID    uuid.UUID
	PassedStatus string
	RegistStatus string

	Subject InternSubject
}
