package main

import "fmt"

func main() {

	// var arr []int = []int{4, 5, 7, 6, 9}
	// var evenNumbers = func(nums ...int) (result []int) {
	// 	for _, num := range nums {
	// 		if num%2 == 0 {
	// 			result = append(result, num)
	// 		}
	// 	}

	// 	return
	// }(arr...)
	// fmt.Println(evenNumbers)

	// students := []string{"Hacktiv8", "Koinworks", "Noobeeid"}

	// find := findStudent(students, "Hacktiv8", checkIsSame)

	// fmt.Println(find)
	max := 50
	fmt.Println("Maximum number :", max)
	fmt.Println("Bilangan Prima :", findPrime(max, func(i int) (isPrime bool) {
		count := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				count++
			}
		}

		if count < 2 {
			isPrime = true
		}

		return
	}))
}

func checkIsSame(str1, str2 string) bool {
	return str1 == str2
}

func findStudent(students []string, search string, cb func(string, string) bool) (res string) {
	isExist := false
	for _, st := range students {
		if cb(st, search) {
			isExist = true
			break
		}
	}

	if !isExist {
		res = fmt.Sprintf("%s not found", search)
	} else {
		res = fmt.Sprintf("%s found in this data", search)
	}

	return
}

func findPrime(max int, cb func(int) bool) (res []int) {

	for i := 1; i <= max; i++ {
		if cb(i) {
			res = append(res, i)
		}
	}

	return
}
