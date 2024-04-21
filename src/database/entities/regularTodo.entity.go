package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegularTodo struct {
	gorm.Model
	StudentID  uuid.UUID
	BusinessID uuid.UUID

	Student  Student
	Business Business
}
