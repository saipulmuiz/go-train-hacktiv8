package params

import (
	"final-project/models"
	"time"
)

type PhotoPostResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"created_at"`
}

type PhotoGetResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      models.UserPhoto
}
