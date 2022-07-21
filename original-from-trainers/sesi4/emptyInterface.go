package main

import "fmt"

func main() {
	var randomValue interface{}

	randomValue = "Hello"
	fmt.Printf("%v\n", randomValue)

	randomValue = "90"
	randInt, ok := randomValue.(int)
	if ok {
		fmt.Println(randInt + 10)
	}

	randomValue = false
	fmt.Printf("%v\n", randomValue)

	randomValue = []int{1, 2, 3}
	fmt.Println(randomValue.([]int)[0])

	randomValue = struct {
		Id   string
		Name string
	}{
		Id:   "1",
		Name: "Hacktiv8",
	}
	fmt.Println(randomValue.(struct {
		Id   string
		Name string
	}).Name)
}
