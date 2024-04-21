package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoAppreciation struct {
	gorm.Model
	RegularTodoDetailID uuid.UUID
	Content             string `gorm:"type:varchar(1000)"`

	RegularTodoDetail RegularTodoDetail
}
