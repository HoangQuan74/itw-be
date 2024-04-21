package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AppliedHistory struct {
	gorm.Model
	JobID              uuid.UUID
	StudentID          uuid.UUID
	CvFileUrl          string `gorm:"type:varchar(1000)"`
	IntroducingFileUrl string `gorm:"type:varchar(1000)"`
	RatePoint          uint   `gorm:"type:smallint"`
	ApplyStatus        string
	DealSalary         float64 `gorm:"type:decimal(20,2)"`

	Job     Job
	Student Student
}
