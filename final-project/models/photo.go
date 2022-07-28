package models

import "time"

type Photo struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption"`
	PhotoUrl  int       `gorm:"not null" json:"photo_url" valid:"required~PhotoUrl is required"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
}
