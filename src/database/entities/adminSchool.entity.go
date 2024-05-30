package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminSchool struct {
	gorm.Model
	AccountStatus int
	UserID        uuid.UUID `gorm:"type:uuid;size:255"`
	SchoolID      uuid.UUID `gorm:"type:uuid;size:255"`

	User   UserPerson
	School School
}
