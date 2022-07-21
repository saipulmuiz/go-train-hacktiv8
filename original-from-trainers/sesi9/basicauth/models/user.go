package models

type User struct {
	Id    int8
	Name  string
	Grade int8
}

var users []User = []User{
	{
		Id:    1,
		Name:  "Hacktiv8",
		Grade: 1,
	},
	{
		Id:    2,
		Name:  "Koinworks",
		Grade: 2,
	},
	{
		Id:    3,
		Name:  "NooBeeID",
		Grade: 3,
	},
}

func GetUsers() *[]User {
	return &users
}

func GetUserById(id int8) *User {
	var u User
	for _, us := range users {
		if us.Id == id {
			u = us
			break
		}
	}

	return &u
}
