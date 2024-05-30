package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Major struct {
	gorm.Model
	Name         string    `gorm:"type:varchar(50)"`
	DepartmentID uuid.UUID `gorm:"type:uuid;size:255"`

	Department Department
}
