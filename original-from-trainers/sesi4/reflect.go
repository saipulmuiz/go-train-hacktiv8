package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 30
	reflectVal := reflect.ValueOf(num)

	fmt.Println("Type variable", reflectVal.Type())

	if reflectVal.Kind() == reflect.Int {
		val := reflectVal.Interface()
		fmt.Println("Nilainya adalah ", val.(int)+1)
	}

	s := &Student{
		Name: "hacktiv8",
	}

	fmt.Println("before :", s.Name)

	ref := reflect.ValueOf(s)
	method := ref.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Koinworks"),
	})

	fmt.Println("after :", s.Name)
}

type Student struct {
	Name string
}

func (s *Student) SetName(name string) {
	s.Name = name
}
