package params

import "time"

type UserRegisterResponse struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type UserUpdateResponse struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
