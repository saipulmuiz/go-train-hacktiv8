package main

import (
	"fmt"
)

func main() {
	defer print("defer 1")
	defer print("defer 2")
	defer print("defer 3")
	flag := false // dapat dari sebuah prosess

	print("Kedua")
	for i := 0; i < 5; i++ {
		print(fmt.Sprintf("%d", i))
	}

	if flag {
		fmt.Println(flag)
		return
	} else {
		fmt.Println("Yah gagal")
		// os.Exit(2)
		return
	}
}

func print(str string) {
	// defer fmt.Println("function print called success")
	fmt.Println(str)
}
