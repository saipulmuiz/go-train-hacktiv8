package main

import "fmt"

func main() {
	// -- Variable with type data --
	// var name string
	// var age int = 23

	// -- Reassign value on variable --
	// name = "Airell"
	// age = 25
	// fmt.Println("Ini adalah namanya ==>", name)
	// fmt.Println("Ini adalah umurnya ==>", age)

	// -- Variable without type data --
	// var name = "Airell"
	// var age = 23

	// -- Short Declaration --
	// name := "Airell"
	// age := 23

	// Print format -> show type data variable

	// -- Multiple variable declarations same data type --
	// var student1, student2, student3 string = "satu", "dua", "tiga"

	// var first, second, third int

	// first, second, third = 1, 2, 3
	// fmt.Println(student1, student2, student3)
	// fmt.Println(first, second, third)

	// -- Multiple variable declarations different data type --
	// var name, age, address = "Airell", 23, "Jalan sudirman"

	// first, second, third := "1", 2, "3"

	// fmt.Println(name, age, address)

	// fmt.Println(first, second, third)

	// --  Underscore Variable (variable yang boleh menganggur) --
	// var firstVariable string

	// var name, age, address = "Airell", 23, "Jalan Sudirman"

	// _, _, _, _ = firstVariable, name, age, address

	// -- fmt.printf Usage for format --

	// first, second := 1, "2"

	// fmt.Printf("Tipe data variable first adalah %T \n", first)

	// fmt.Printf("Tipe data variable second adalah %T \n", second)

	//  -- fmt.printf Usage Another
	var name = "Airell"

	var age = 23

	var address = "Jalan Sudirman"

	fmt.Printf("Halo nama ku %s, umur aku adalah %d, dan aku tinggal di %s", name, age, address)
}
