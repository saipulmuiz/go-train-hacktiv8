package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums [4]int = [4]int{1, 2, 3, 4}
	nums[2] = 10

	for _, n := range nums {
		fmt.Println(n)
	}

	var mhs [3][2]string = [3][2]string{
		{"1", "mhs1"},
		{"2", "mhs2"},
		{"3", "mhs3"},
	}

	for _, mhsNum := range mhs {
		for j, m := range mhsNum {
			if j == 1 {
				fmt.Printf("Nama : %s\n", m)
			} else {
				fmt.Printf("ID : %s\n", m)
			}
		}
		fmt.Println(strings.Repeat("=", 20))
	}
}
