package main

import "fmt"

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

	var products = []map[string]string{
		{
			"id":    "1",
			"name":  "product1",
			"stock": "3",
		},
		{
			"id":    "2",
			"name":  "product2",
			"stock": "2",
		},
		{
			"id":    "1",
			"name":  "product3",
			"stock": "1",
		},
	}

	var trxs = []map[string]string{
		{
			"id":         "1",
			"username":   "admin",
			"id_product": "1",
			"amount":     "3000",
		},
	}

	for _, trx := range trxs {
		fmt.Printf("Id: %s, Username: %s, Id Product: %s, Amount: %s\n", trx["id"], trx["username"], trx["id_product"], trx["amount"])
	}
	for _, product := range products {
		fmt.Printf("Id: %s, name: %s, stock: %s\n", product["id"], product["name"], product["stock"])
	}

}
