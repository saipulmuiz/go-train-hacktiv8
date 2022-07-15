package main

import (
	"fmt"
)

func main() {
	// -- Deklarasi pointer --
	// var firstNumber *int
	// var secondNumber *int

	// _, _ = firstNumber, secondNumber

	// -- Pointer (Memory Address) --
	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber

	// fmt.Println("firstNumber (value) :", firstNumber)
	// fmt.Println("firstNumber (Memory address) :", &firstNumber)

	// fmt.Println("secondNumber (value) :", *secondNumber)
	// fmt.Println("secondNumber (Memory address) :", &secondNumber)

	// -- Pointer (Changing value through pointer) --
	// var firstPerson string = "John"
	// var secondPerson *string = &firstPerson

	// fmt.Println("firstPerson (value) :", firstPerson)
	// fmt.Println("firstPerson (memory address) :", &firstPerson)
	// fmt.Println("secondPerson (value) :", *secondPerson)
	// fmt.Println("secondPerson (memory address) :", secondPerson)

	// fmt.Println("\n", strings.Repeat("#", 30), "\n")

	// *secondPerson = "Doe"

	// fmt.Println("firstPerson (value) :", firstPerson)
	// fmt.Println("firstPerson (memory address) :", &firstPerson)
	// fmt.Println("secondPerson (value) :", *secondPerson)
	// fmt.Println("secondPerson (memory address) :", secondPerson)

	// -- Pointer (Pointer as a parameter) --
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
