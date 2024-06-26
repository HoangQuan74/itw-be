package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name             string    `gorm:"type:varchar(50)"`
	DepartmentHeadID uuid.UUID `gorm:"type:uuid"`
	SchoolID         uuid.UUID `gorm:"type:uuid"`

	DepartmentHead *Teacher `gorm:"foreignKey:DepartmentHeadID"`
	School         School
}
