package main

import (
	"fmt"
	"strings"
)

func main() {
	var arr = [4]string{"Mhs 1", "Mhs 2", "Mhs 4", "Mhs 5"}
	var sl = arr[:3]
	fmt.Println("arr", arr)
	fmt.Println("slice", sl)

	var s []string
	s = append(s, sl...)

	sl[0] = "Mhs 10"
	fmt.Println("arr", arr)
	fmt.Println("slice", sl)

	fmt.Println(s)

	var students = []string{"Mhs 1", "Mhs 2"}

	fmt.Println("Panjang :", len(students))
	fmt.Println("Data :", students)
	fmt.Println("Kapasitas :", cap(students))

	students = append(students, "Mhs 3")
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Penambahan 1 data")
	fmt.Println("Panjang :", len(students))
	fmt.Println("Data :", students)
	fmt.Println("Kapasitas :", cap(students))

	var newS = []string{"Mhs 4", "Mhs 5", "Mhs 6", "Mhs 7", "Mhs 8", "Mhs 9"}
	students = append(students, newS...)
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Penambahan 6 data")
	fmt.Println("Panjang :", len(students))
	fmt.Println("Data :", students)
	fmt.Println("Kapasitas :", cap(students))

	students = append(students, "Mhs 10")
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Penambahan 1 data")
	fmt.Println("Panjang :", len(students))
	fmt.Println("Data :", students)
	fmt.Println("Kapasitas :", cap(students))

	students = append(students, "Mhs 11")
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Penambahan 1 data")
	fmt.Println("Panjang :", len(students))
	fmt.Println("Data :", students)
	fmt.Println("Kapasitas :", cap(students))
}
