package entities

import "github.com/google/uuid"

type UserAccount struct {
	UserName string `gorm:"primaryKey;type:varchar(50)"`
	Pass     string `gorm:"type:varchar(300)"`
	RoleID   uuid.UUID `gorm:"size:255"`

	Role Role
}
