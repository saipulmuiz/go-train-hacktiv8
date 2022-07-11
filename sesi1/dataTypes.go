package main

import "fmt"

func main() {
	// -- Data type integer --
	// var first = 89
	// var second = -12

	// -- Data type integer uint and int --
	// var first uint8 = 89
	// var second int8 = -12

	// fmt.Printf("Tipe data first : %T \n", first)
	// fmt.Printf("Tipe data second : %T \n", second)

	// -- data type float --
	// var decimalNumber float32 = 3.63

	// fmt.Printf("decimal Number: %f\n", decimalNumber)
	// fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	// -- data type bool
	// var condition bool = true
	// fmt.Printf("is it permitted? %t \n", condition)

	// -- data type string --
	// var message string = "Halo"
	var message = `Selamat datang di "Hactiv8".
	Salam kenal :).
	Mari belajar Scalable Web Service With Go".`
	fmt.Print(message)
}
