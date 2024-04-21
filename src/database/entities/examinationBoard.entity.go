package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExaminationBoard struct {
	gorm.Model
	PresidentID    uuid.UUID
	SecretaryID    uuid.UUID
	AskerID        uuid.UUID
	AcademicYearID uuid.UUID
	SemesterID     uuid.UUID
	DepartmentID   uuid.UUID
	ReportDate     *time.Time

	President    Teacher
	Secretary    Teacher
	Asker        Teacher
	AcademicYear AcademicYear
	Semester     Semester
	Department   Department
}
