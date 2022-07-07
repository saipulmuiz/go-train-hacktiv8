package main

import "fmt"

func main() {
	var nums [4]int = [4]int{1, 2, 3}
	fmt.Println(nums)

	for _, n := range nums {
		fmt.Println(n)
	}
}
