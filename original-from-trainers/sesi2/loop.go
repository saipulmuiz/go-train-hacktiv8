package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	counter := 0
	for counter < 5 {
		fmt.Println("counter", counter+1)
		counter++
	}

	for {
		if counter%2 == 0 {
			counter--
			break
		}
		fmt.Println("counter ke2", counter)
		counter--
		if counter == 0 {
			break
		}
	}

bebas:
	for i := 0; i < 6; i++ {
		fmt.Println("Perulangan ke", i)
	coba:
		for j := 0; j < 3; j++ {
			if i == 2 {
				break bebas
			} else {
				break coba
				// fmt.Print(j, " ")
			}
		}
	}
}
