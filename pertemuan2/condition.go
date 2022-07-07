package main

import "fmt"

func main() {
	var currentYear = 2022

	// if currentYear == 2022 {
	// 	fmt.Println("Sekarang tahun dua ribu dua puluh dua")
	// } else {
	// 	fmt.Println("Sekarang tahun berapa yaa??")
	// }

	switch currentYear {
	case 2022:
		fmt.Println("Sekarang tahun dua ribu dua puluh dua")
	default:
		fmt.Println("Sekarang tahun berapa yaa??")
	}
}
