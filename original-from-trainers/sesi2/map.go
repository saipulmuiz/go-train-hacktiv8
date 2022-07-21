package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var mhs1 = map[string]string{
		"name": "Mhs 1",
		"id":   "1",
	}

	var mhs = []map[string]string{
		{
			"name": "Mhs 1",
			"id":   "1",
		},
		{
			"name": "Mhs 2",
			"id":   "2",
		},
		{
			"name": "Mhs 3",
			"id":   "3",
		},
		{
			"name": "Mhs 4",
			"id":   "4",
		},
	}

	mhs1["age"] = "20"

	umur, isExist := mhs1["age"]

	if isExist {
		umurInt, err := strconv.Atoi(umur)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(umurInt)
		}
	}

	for _, m := range mhs {
		for key, value := range m {
			fmt.Printf("%s : %s \n", key, value)
		}
		fmt.Println(strings.Repeat("=", 20))
	}

}
