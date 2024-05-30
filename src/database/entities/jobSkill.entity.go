package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobSkill struct {
	gorm.Model
	JobID   uuid.UUID `gorm:"size:255"`
	SkillID uuid.UUID `gorm:"size:255"`

	Job   Job
	Skill Skill
}
