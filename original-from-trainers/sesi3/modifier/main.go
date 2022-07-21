package main

import (
	"fmt"
	"modifier/constants"
	"modifier/helper"
	"modifier/models"
	"os"
)

func main() {
	args := os.Args

	curr := constants.CurrencyName("IDR")
	fmt.Println(args[1], curr)

	emp := models.Employee{
		Name: "Emp1",
		Age:  10,
	}

	helper.Calculate()

	fmt.Println(emp, helper.Res)
}
