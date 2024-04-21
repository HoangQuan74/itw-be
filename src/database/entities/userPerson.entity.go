package entities

import (
	"gorm.io/gorm"
)

type UserPerson struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50)"`
	FullName string `gorm:"type:varchar(100)"`
	Image    string `gorm:"type:varchar(1000)"`
	Phone    string `gorm:"type:varchar(11)"`
	Email    string `gorm:"type:varchar(50)"`
	Address  string `gorm:"type:varchar(200)"`

	UserAccount UserAccount `gorm:"foreignKey:UserName"`
}
