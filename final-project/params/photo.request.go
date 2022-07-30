package params

type PhotoRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo Url is required"`
	UserId   uint   `json:"user_id"`
}
