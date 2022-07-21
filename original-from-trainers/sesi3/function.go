package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{10, 20, 50, 20}
	res := calculate("+", nums...)
	fmt.Println(res)
}

func greet(name, age string) (result string) {
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return
	}
	result = fmt.Sprintf("Hello %s, saya berumur %d", name, ageInt)
	return
}

func calculate(operator string, nums ...int) (total int) {
	total = nums[0]

	for i := 1; i < len(nums); i++ {
		switch operator {
		case "+":
			total += nums[i]
		case "-":
			total -= nums[i]
		}
	}
	return
}
