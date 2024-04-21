package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobSkill struct {
	gorm.Model
	JobID   uuid.UUID
	SkillID uuid.UUID

	Job   Job
	Skill Skill
}
