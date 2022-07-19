package models

type Product struct {
	GormModel
	Title       string
	Description string
	UserID      uint
	User        *User
}
