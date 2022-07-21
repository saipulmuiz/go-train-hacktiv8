package main

import "fmt"

var age = 20

func main() {
	age := 10
	if true {
		age := 3
		fmt.Println(age)
	}
	fmt.Println(age)
	test()
}

func test() {
	fmt.Println(age)
}
