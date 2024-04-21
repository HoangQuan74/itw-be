package entities

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	EstablishDate  *time.Time
	IndustrySector string `gorm:"type:varchar(50)"`
	Representator  string `gorm:"type:varchar(50)"`
	ShortDesc      string `gorm:"type:varchar(1000)"`
}
