package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		2 slice of map
			- slice of products
				- id
				- name
				- stock
			- slice of trx
				- id
				- username
				- id products
				- amount
	*/

	products := []map[string]string{
		{
			"id":    "1",
			"name":  "Baju",
			"stock": "20",
		},
		{
			"id":    "2",
			"name":  "Celana",
			"stock": "2",
		},
		{
			"id":    "1",
			"name":  "Meja",
			"stock": "100",
		},
	}

	trxs := []map[string]string{
		{
			"id":         "1",
			"username":   "Hacktiv8",
			"amount":     "20",
			"id_product": "1",
		},
		{
			"id":         "2",
			"username":   "Koinworks",
			"amount":     "20",
			"id_product": "2",
		},
		{
			"id":         "1",
			"username":   "Noobee",
			"amount":     "20",
			"id_product": "4",
		},
	}

	for _, t := range products {
		pName := ""
		pStock := ""
		isExist := false
		for _, p := range products {
			if t["id_products"] == p["id"] {
				pName = p["name"]
				pStock = p["stock"]
				isExist = true
				break
			}
		}

		fmt.Println("username :", t["username"])
		fmt.Println("id :", t["id"])
		fmt.Println("amount :", t["amount"])
		if isExist {
			fmt.Println("product name :", pName)
			fmt.Println("product stock :", pStock)
		} else {
			fmt.Println("id product :", t["id_product"])
		}

		fmt.Println(strings.Repeat("=", 20))
	}

	for _, trx := range trxs {
		fmt.Printf("Id: %s, Username: %s, Id Product: %s, Amount: %s\n", trx["id"], trx["username"], trx["id_product"], trx["amount"])
	}
	for _, product := range products {
		fmt.Printf("Id: %s, name: %s, stock: %s\n", product["id"], product["name"], product["stock"])
	}

}
