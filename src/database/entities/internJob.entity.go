package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternJob struct {
	gorm.Model
	StudentID           uuid.UUID `gorm:"size:255"`
	StartDate           *time.Time
	EndDate             *time.Time
	ApplyID             uuid.UUID `gorm:"size:255"`
	IsInterning         string
	AppreciationFileUrl string `gorm:"type:varchar(1000)"`

	Student Student
	apply   AppliedHistory
}
