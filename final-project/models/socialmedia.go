package models

import "time"

type SocialMedia struct {
	Id             uint      `gorm:"primary_key" json:"id"`
	Name           string    `gorm:"not null" json:"name" valid:"required~Name is required"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url" valid:"required~Social Media URL is required"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User
}
