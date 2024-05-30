package student

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

type StudentInput struct {
	AdmissionDate *time.Time `json:"admission_date" validate:"required"`
	Dob           *time.Time `json:"dob" validate:"required"`
	Sex           string     `json:"sex" validate:"required,oneof=male female"`
	CurrentStatus int        `json:"current_status" validate:"required"`
	ClassID       uuid.UUID  `json:"class_id" validate:"required"`
}
