package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	StudentID             uuid.UUID
	ReportFileUrl         string `gorm:"type:varchar(100)"`
	ResultBusinessFileUrl string `gorm:"type:varchar(1000)"`
	ResultTeacherFileUrl  string `gorm:"type:varchar(1000)"`
	SubjectID             uuid.UUID

	Student Student
	Subject InternSubject
}
