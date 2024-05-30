package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Name              string    `gorm:"type:varchar(50)"`
	Requirements      string    `gorm:"type:varchar(1000)"`
	AnotherInfomation string    `gorm:"type:varchar(1000)"`
	Vacancies         int       `gorm:"type:smallint"`
	BusinessID        uuid.UUID `gorm:"size:255"`
	StandardSalary    float64   `gorm:"type:decimal(20,2)"`
	AverageRate       float32   `gorm:"type:decimal(3,2)"`
	ViewerCount       int
	WorkType          string
	WorkSpace         string
	ExpiredDate       *time.Time
	ExperienceYear    int       `gorm:"type:smallint"`
	PositionID        uuid.UUID `gorm:"size:255"`
	ImageUrl          string    `gorm:"type:varchar(1000)"`
	JobDesc           string    `gorm:"type:varchar(1000)"`
	LevelID           uuid.UUID `gorm:"size:255"`

	Business Business
	Position Position
	Level    JobLevel
}
