package params

import "time"

type PhotoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  int       `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
