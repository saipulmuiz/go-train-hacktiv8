package params

type SocialMediaRequest struct {
	Name           string `json:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social Media URL is required"`
	UserId         uint   `json:"user_id"`
}
