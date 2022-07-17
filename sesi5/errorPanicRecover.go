package main

import (
	"errors"
	"fmt"
)

func main() {
	// -- Error --
	// var number int
	// var err error

	// number, err = strconv.Atoi("123GH")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// number, err = strconv.Atoi("123")

	// if err == nil {
	// 	fmt.Println(number)
	// } else {
	// 	fmt.Println(err.Error())
	// }

	// -- Error (Custom error) --
	// var password string

	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	// -- Panic --
	// var password string

	// fmt.Scanln(&password)

	// if valid, err := validPassword(password); err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Println(valid)
	// }

	// -- Recover --
	defer catchErr()
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

// --Recover --
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("pasword has to have more than 4 characters")
	}

	return "Valid password", nil
}
