package main

import "exportimport/helpers"

func main() {
	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()
}
