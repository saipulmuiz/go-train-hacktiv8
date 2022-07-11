package main

import "fmt"

func main() {
	fmt.Println(isCanWork("female", 21))
}

func isCanWork(gender string, age int) string {
	if gender == "male" && age > 17 {
		return "Gender male boleh bekerja"
	} else if gender == "female" && age > 20 {
		return "Gender female boleh bekerja"
	} else {
		return "gak boleh bekerja"
	}
}
