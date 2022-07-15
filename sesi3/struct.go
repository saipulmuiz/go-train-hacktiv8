package main

import (
	"fmt"
)

// -- Struct (Embeded Struct) --\
type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

// -- Penulisan struct --
// type Employee struct {
// 	name     string
// 	age      int
// 	division string
// }

func main() {
	// -- Struct (Giving value to struct) --
	// var employee Employee

	// employee.name = "Airell"

	// employee.age = 23

	// employee.division = "Curriculum Developer"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)
	// fmt.Println(employee.division)

	// -- Struct (Initializing struct) --
	// var employee1 = Employee{}
	// employee1.name = "Airell"
	// employee1.age = 23
	// employee1.division = "Curriculum Developer"

	// var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}

	// fmt.Printf("Employee1: %+v\n", employee1)
	// fmt.Printf("Employee2: %+v\n", employee2)

	// -- Struct (Pointer to a struct) --
	// var employee1 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}

	// var employee2 *Employee = &employee1

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// fmt.Println(strings.Repeat("#", 20))

	// employee2.name = "Ananda"

	// fmt.Println("Employee1 name:", employee1.name)
	// fmt.Println("Employee2 name:", employee2.name)

	// -- Struct (Embeded Struct) --

	// var employee1 = Employee{}

	// employee1.person.name = "Airell"
	// employee1.person.age = 23
	// employee1.division = "Curriculum Developer"

	// fmt.Printf("%+v", employee1)

	// -- Struct (Anonymous struct) --

	// Anonymous struct tanpa pengisian field
	// var employee1 = struct {
	// 	person   Person
	// 	division string
	// }{}
	// employee1.person = Person{name: "Airell", age: 23}
	// employee1.division = "Curriculum Developer"

	// // Anonymous struct dengan pengisisan field
	// var employee2 = struct {
	// 	person   Person
	// 	division string
	// }{
	// 	person:   Person{name: "Ananda", age: 23},
	// 	division: "Finance",
	// }

	// fmt.Printf("Employee 1 : %+v\n", employee1)
	// fmt.Printf("Employee 2 : %+v\n", employee2)

	// -- Struct (Slice of struct) --
	// var people = []Person{
	// 	{name: "Airell", age: 23},
	// 	{name: "Ananda", age: 23},
	// 	{name: "Mailo", age: 23},
	// }

	// for _, v := range people {
	// 	fmt.Printf("%+v\n", v)
	// }

	// -- Struct (Slice of anonymous struct) --
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Airell", age: 23}, division: "Curriculum Developer"},
		{person: Person{name: "Ananda", age: 23}, division: "Finance"},
		{person: Person{name: "Mailo", age: 21}, division: "Marketing"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
