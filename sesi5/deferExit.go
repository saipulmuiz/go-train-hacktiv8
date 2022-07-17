package main

import (
	"fmt"
	"os"
)

func main() {
	// -- Defer Implementation --
	// defer fmt.Println("defer function starts to execute") // Dijalan kan diakhir setelah semua perintah beres
	// fmt.Println("Hai everyone")
	// fmt.Println("Welcome back ti Go Learning Center")

	// callDeferFunc()
	// fmt.Println("Hai everyone!")

	// -- Exit Implementation --
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
