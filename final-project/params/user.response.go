package params

import "time"

type UserRegisterResponse struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
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
