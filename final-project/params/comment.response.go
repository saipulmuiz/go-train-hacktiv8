package params

import (
	"final-project/models"
	"time"
)

type CommentPostResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentUpdateResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentGetResponse struct {
	Id        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   uint      `json:"photo_id"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      models.UserComment
	Photo     models.PhotoComment
}
