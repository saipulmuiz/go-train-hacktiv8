package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Status  string
	payload []Employee
	Query   Query
}

type Employee struct {
	Id    string
	Name  string
	Email string
}

type Query struct {
	Limit      int
	Page       int
	Total_page int
}

func main() {
	jsonData := `
	{
		"status" : "success",
		"payload" : [
			{
				"id": "1",
				"name": "Hactiv8",
				"email": "admin@hacktiv8.com"
			},
			{
				"id": "2",
				"name": "Koinworks",
				"email": "admin@koinworks.com"
			}
		],
		"query" : {
			"limit" : 20,
			"page" : 1,
			"total_page" : 5
		}
	}
	`

	// var emp map[string]interface{}
	var emp Data
	err := json.Unmarshal([]byte(jsonData), &emp)
	if err != nil {
		panic(err.Error())
	}

	// payload := emp.payload

	// payload := emp["payload"]

	// payloadByte, err := json.Marshal(payload)

	// var e Employee
	// err = json.Unmarshal(payloadByte, &e)
	// for _, v := range payload {
	// 	fmt.Printf("Payload : %+v\n", v)
	// }

	fmt.Printf("%+v", emp)
}
