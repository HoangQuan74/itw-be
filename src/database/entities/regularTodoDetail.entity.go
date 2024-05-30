package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegularTodoDetail struct {
	gorm.Model
	Name            string `gorm:"type:varchar(200)"`
	StartDate       *time.Time
	EndDate         *time.Time
	RegularTodoID   uuid.UUID `gorm:"size:255"`
	CurrentStatus   string
	CompletedStatus string

	RegularTodo RegularTodo
}
