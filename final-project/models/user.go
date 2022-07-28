package models

import (
	"final-project/helper"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null;uniqueIndex;unique" json:"username" valid:"required~Username is required"`
	Email     string    `gorm:"not null;uniqueIndex;unique" json:"email" valid:"required~Email is required, email~Invalid format email"`
	Password  string    `gorm:"not null" json:"password" valid:"required~Password is required, minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age       uint      `gorm:"not null" json:"age" valid:"required~Age is required, range(8|80)~Age has to have minimum value above of 8 to 80"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return
}
