package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminBusiness struct {
	gorm.Model
	UserID     uuid.UUID
	BusinessID uuid.UUID

	User     UserPerson
	Business Business
}
