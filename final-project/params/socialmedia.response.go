package params

import (
	"final-project/models"
	"time"
)

type SocialMediaPostResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaGetResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           models.UserSocialMedia
}
