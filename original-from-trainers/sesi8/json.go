package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Status  string
	Payload []Employee
}

type Employee struct {
	Id    string
	Name  string
	Email string
}

func main() {
	jsonData := `
		{
			"status" : "success",
			"payload" : [
				{
					"id" : "1",
					"name" : "Hacktiv8",
					"email" : "admin@hacktiv8.com"
				},
				{
					"id" : "1",
					"name" : "Hacktiv8",
					"email" : "admin@hacktiv8.com"
				}
			],
			"query" : {
				"limit" : 20,
				"page" : 1,
				"total_page" : 5
			}
		}
	`

	var emp Data

	err := json.Unmarshal([]byte(jsonData), &emp)
	if err != nil {
		panic(err.Error())
	}

	// payload := emp["payload"]

	// payloadByte, err := json.Marshal(payload)

	// var e Employee
	// err = json.Unmarshal(payloadByte, &e)

	fmt.Printf("%+v", emp)
}
