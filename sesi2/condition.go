package main

import "fmt"

func main() {
	// -- If Condition (temporary variable) --
	// var currentYear = 2022

	// if age := currentYear - 1998; age < 17 {
	// 	fmt.Println("Kamu belum boleh membuat kartu sim")
	// } else {
	// 	fmt.Println("Kamu sudah boleh membuat kartu sim")
	// }

	// -- Switch --
	// var score = 8

	// switch score {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7:
	// 	fmt.Println("awesome")
	// default:
	// 	fmt.Println("not bad")
	// }

	// -- Switch with relational operators --
	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// default:
	// 	{
	// 		fmt.Println("study harder")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// -- Switch (fallthrough keyword) --
	// var score = 6

	// switch {
	// case score == 8:
	// 	fmt.Println("perfect")
	// case (score < 8) && (score > 3):
	// 	fmt.Println("not bad")
	// 	fallthrough
	// case score < 5:
	// 	fmt.Println("It is ok, but please study harder")
	// default:
	// 	{
	// 		println("study harder")
	// 		println("You don't have a good score yet")
	// 	}
	// }

	// -- Nested Conditions --
	var score = 10

	if score > 17 {
		switch score {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
