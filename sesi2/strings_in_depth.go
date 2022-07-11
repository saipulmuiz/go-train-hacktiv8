package main

import "fmt"

func main() {
	// city := "Jakarta"

	// for i := 0; i < len(city); i++ {
	// 	fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	// }

	var city []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city))
}
