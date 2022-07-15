package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// var names = []string{"Airell", "Jordan"}
	// var printMsg = greetReturn("Heii", names)

	// fmt.Println(printMsg)

	// var diameter float64 = 15

	// var area, circumference = calculate2(diameter)

	// fmt.Println("Area:", area)
	// fmt.Println("Circumference:", circumference)

	studentLists := print("Airell", "Nanda", "Mailo", "Schannel", "Marco")

	fmt.Printf("%v", studentLists)
}

func greet(name string, age int) {
	fmt.Println("Hello", name, "\n Saya berumur", age)
}

// -- Function return value --
func greetReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// -- Function returning multiple values --
func calculate(d float64) (float64, float64) {
	// Menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)
	// Menghitung keliling

	var circumference = math.Pi * d

	return area, circumference
}

// -- Function (Predefined return value) --
func calculate2(d float64) (area float64, circumference float64) {
	// Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	// Menghitung keliling
	circumference = math.Pi * d

	return
}

// -- Function Variadic --
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}
