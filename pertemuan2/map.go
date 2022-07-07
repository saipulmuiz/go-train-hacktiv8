package main

import "fmt"

func main() {
	var mhs1 = map[string]string{
		"id":   "1",
		"name": "Mhs 1",
	}

	fmt.Println("Mhs 1", mhs1)
	fmt.Println("Nama :", mhs1["name"])
}
