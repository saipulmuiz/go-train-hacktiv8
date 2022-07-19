package models

type User struct {
	GormModel
	FullName string
	Email    string
	Password string
	Products []Product
}
