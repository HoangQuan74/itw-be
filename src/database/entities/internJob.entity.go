package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternJob struct {
	gorm.Model
	StudentID           uuid.UUID
	StartDate           *time.Time
	EndDate             *time.Time
	ApplyID             uuid.UUID
	IsInterning         string
	AppreciationFileUrl string `gorm:"type:varchar(1000)"`

	Student Student
	apply   AppliedHistory
}
