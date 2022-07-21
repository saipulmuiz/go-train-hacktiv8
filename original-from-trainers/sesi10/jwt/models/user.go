package models

import (
	"jwt/helper"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" valid:"required~Full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required, email~Invalid format email"`
	Password string    `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// _, errCreate := govalidator.ValidateStruct(u)
	// if errCreate != nil {
	// 	return errCreate
	// }

	hash, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return
}
