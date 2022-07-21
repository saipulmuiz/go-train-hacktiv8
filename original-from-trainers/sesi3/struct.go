package main

import "fmt"

type Employee struct {
	Name        string
	Age         int
	Division    Division
	SocialMedia []string
	Dob         int
}

type Division struct {
	Name    string
	Manager string
}

func main() {
	var emps []*Employee

	emps = []*Employee{
		{
			Name: "Hacktiv8",
			Division: Division{
				Name:    "IT",
				Manager: "Admin",
			},
		},
	}
	employee := Employee{
		Name: "Hacktiv8",
		Dob:  1998,
		Division: Division{
			Name:    "IT",
			Manager: "Koinworks",
		},
		SocialMedia: []string{
			"Instagram", "WA",
		},
	}

	emps = append(emps, &employee)
	print(&emps)
	for _, e := range emps {
		fmt.Printf("%+v \n", *e)
	}
}

func print(emps *[]*Employee) {
	for _, e := range *emps {
		e.changeName("Noobeeid")
		e.Division.changeName("Ecommerce")
		fmt.Printf("%+v \n", e)
	}
}

func (e *Employee) changeName(name string) {
	e.Name = name
}

func (e *Division) changeName(name string) {
	e.Name = name
}

func changeName(emp *Employee, name string) {
	emp.Name = name
	age := 2022 - 1998

	emp.Age = age
}
