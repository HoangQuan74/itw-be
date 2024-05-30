package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternSubject struct {
	gorm.Model
	Name           string    `gorm:"type:varchar(255)"`
	Unit           int       `gorm:"type:smallint"`
	Sessions       int       `gorm:"type:smallint"`
	MaxStudent     int       `gorm:"type:smallint"`
	TeacherID      uuid.UUID `gorm:"size:255"`
	AcademicYearID uuid.UUID `gorm:"size:255"`
	SemesterID     uuid.UUID `gorm:"size:255"`
	StartDate      *time.Time
	EndDate        *time.Time
	IsOpen         int `gorm:"type:tinyint"`

	Teacher      Teacher
	AcademicYear AcademicYear
	Semester     Semester
}
