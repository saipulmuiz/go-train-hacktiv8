package main

import (
	"fmt"
	"strings"
)

func main() {
	number := 10
	number2 := number
	fmt.Println(number, &number)
	check(number)
	fmt.Println(number, &number)
	check(number2)

	change(&number, 30)

	var pointerNumber *int = &number
	fmt.Println(number, &number)
	fmt.Println(*pointerNumber, pointerNumber)

	changeNumber(&number, 20)
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println(number, &number)
	fmt.Println(*pointerNumber, pointerNumber)

	changeNumber(pointerNumber, 100)
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println(number, &number)
	fmt.Println(*pointerNumber, pointerNumber)
}

func change(num *int, into int) {
	*num = into
}

func check(num int) {
	fmt.Println(num, &num)
	test(num)
}

func test(num int) {
	fmt.Println(num, &num)
	fmt.Println("##")
}

func changeNumber(num *int, change int) {
	*num = change
}
