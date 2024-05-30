package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExaminationBoard struct {
	gorm.Model
	PresidentID    uuid.UUID `gorm:"size:255"`
	SecretaryID    uuid.UUID `gorm:"size:255"`
	AskerID        uuid.UUID `gorm:"size:255"`
	AcademicYearID uuid.UUID `gorm:"size:255"`
	SemesterID     uuid.UUID `gorm:"size:255"`
	DepartmentID   uuid.UUID `gorm:"size:255"`
	ReportDate     *time.Time

	President    Teacher
	Secretary    Teacher
	Asker        Teacher
	AcademicYear AcademicYear
	Semester     Semester
	Department   Department
}
