package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminBusiness struct {
	gorm.Model
	UserID     uuid.UUID `gorm:"size:255"`
	BusinessID uuid.UUID `gorm:"size:255"`

	User     UserPerson
	Business Business
}
