package main

import "fmt"

const DB_HOST string = "localhost"
const DB_PORT string = "5432"

func main() {
	const nama string = DB_HOST
	fmt.Println(nama, DB_PORT)
}
