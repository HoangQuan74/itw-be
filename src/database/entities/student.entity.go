package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	UserID        uuid.UUID `gorm:"type:uuid"`
	AdmissionDate *time.Time
	Dob           *time.Time
	Sex           string `gorm:"type:varchar(10)"`
	CurrentStatus int
	ClassID       uuid.UUID `gorm:"type:uuid"`

	User  UserPerson
	Class Class
}
