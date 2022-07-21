package main

import "fmt"

func main() {
	/*
		gender = male || female
		age = int

		isCanWork ?

		1. male, umur > 17
		2. female, umur > 20

		"gender male boleh bekerja ? false"
	*/

	gender := "male"
	age := 18

	isMale := gender == "male"

	isCanWork := (isMale && age > 17) || (!isMale && age > 20)

	fmt.Printf("gender %s boleh bekerja ? %t \n", gender, isCanWork)
}
