package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name           string `gorm:"type:varchar(50)"`
	Students       *int
	AcademicYearID uuid.UUID `gorm:"type:uuid"`
	ProgramID      uuid.UUID `gorm:"type:uuid"`
	MajorID        uuid.UUID `gorm:"type:uuid"`
	TeacherHeadID  uuid.UUID `gorm:"type:uuid"`

	AcademicYear AcademicYear
	Program      Program
	Major        Major
	TeacherHead  Teacher
}
