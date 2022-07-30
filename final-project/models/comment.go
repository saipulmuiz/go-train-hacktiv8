package models

import "time"

type Comment struct {
	Id        uint      `gorm:"primary_key" json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Photo     Photo
}
