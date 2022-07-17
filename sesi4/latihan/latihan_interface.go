package main

import "fmt"

var products = []string{}
var transactions = []string{}

type IProducts interface {
	createProduct() bool
	getProducts() []product
}

type ITransactions interface {
	createTrx() bool
	getTrxById() transaction
}

type product struct {
	id, stok int
	name     string
}

type transaction struct {
	id, amount int
	product    product
}

func (p product) createProduct() bool {
	append(products, product)
	return true
}

func (p product) getProducts() product {
	return p.getProducts()
}

func (t transaction) createTrx() bool {
	return true
}

func (t transaction) getTrxById() int {
	return t.id
}

func main() {
	var product1 IProducts = product{id: 1, name: "product1", stok: 5}

	fmt.Printf("Type of product1: %T", product1)
}
