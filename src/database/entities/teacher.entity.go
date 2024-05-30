package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	UserID         uuid.UUID `gorm:"type:uuid;size:255"`
	Dob            *time.Time
	StartDate      *time.Time
	EducationLevel string `gorm:"type:varchar(30)"`
	ExperienceYear int    `gorm:"type:smallint"`
	CurrentStatus  int
	DepartmentID   uuid.UUID `gorm:"type:uuid;size:255"`

	User       UserPerson
	Department *Department
}
