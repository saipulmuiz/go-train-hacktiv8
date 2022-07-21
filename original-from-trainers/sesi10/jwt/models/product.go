package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Description string `gorm:"not null" json:"description" valid:"required~Description is required"`
	UserID      uint

	User *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		return errCreate
	}

	return
}
