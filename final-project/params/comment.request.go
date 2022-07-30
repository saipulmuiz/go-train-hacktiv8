package params

type CommentRequest struct {
	Message string `json:"message" valid:"required~Message is required"`
	PhotoId uint   `json:"photo_id"`
	UserId  uint   `json:"user_id"`
}
