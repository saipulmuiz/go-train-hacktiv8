package main

import (
	"fmt"
)

func main() {
	// -- Arithmetic Operators --
	// -- Eksekusi perkalian dulu
	// var value = 2 + 2*3
	// fmt.Println(value)

	// -- Ekseskusi pertambahan dulu
	// var value = (2 + 2) * 3
	// fmt.Println(value)

	// -- Relational Operators --
	// var firstCondition bool = 2 < 3
	// var secondCondition bool = "joey" == "Joey"
	// var thirdCondition bool = 10 != 23
	// var fourthCondition bool = 11 <= 11

	// fmt.Println("first condition:", firstCondition)
	// fmt.Println("second condition:", secondCondition)
	// fmt.Println("third condition:", thirdCondition)
	// fmt.Println("fourth condition:", fourthCondition)

	// -- Logical operators --
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
