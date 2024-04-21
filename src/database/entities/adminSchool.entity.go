package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminSchool struct {
	gorm.Model
	AccountStatus int
	UserID        uuid.UUID `gorm:"type:uuid"`
	SchoolID      uuid.UUID `gorm:"type:uuid"`

	User   UserPerson
	School School
}
