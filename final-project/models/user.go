package models

import (
	"final-project/helper"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null;uniqueIndex;unique" json:"username"`
	Email     string    `gorm:"not null;uniqueIndex;unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Age       uint      `gorm:"not null" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserComment struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSocialMedia struct {
	Id              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return
}
