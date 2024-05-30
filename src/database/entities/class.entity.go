package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name           string `gorm:"type:varchar(50)"`
	Students       *int
	AcademicYearID uuid.UUID `gorm:"type:uuid;size:255"`
	ProgramID      uuid.UUID `gorm:"type:uuid;size:255"`
	MajorID        uuid.UUID `gorm:"type:uuid;size:255"`
	TeacherHeadID  uuid.UUID `gorm:"type:uuid;size:255"`

	AcademicYear AcademicYear
	Program      Program
	Major        Major
	TeacherHead  Teacher
}
