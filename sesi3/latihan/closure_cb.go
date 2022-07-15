package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}

	var find = findPrimeNumbers(numbers, func(number int) bool {
		return number%number == 0
	})

	fmt.Println("Prime numbers is =", find)
}

func findPrimeNumbers(numbers []int, cb func(int) bool) []int {
	var result []int

	for _, n := range numbers {
		if cb(n) {
			result = append(result, n)
		}
	}

	return result
}
