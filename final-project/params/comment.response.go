package params

import "time"

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
