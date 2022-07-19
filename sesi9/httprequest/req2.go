package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type Product struct {
	Id    int
	Name  string
	Stock int
}

type ResOJKSingle struct {
	Data  Illegals
	Error interface{}
}

type Illegals struct {
	Version  string
	Illegals Ojk
}

type Ojk struct {
	Id      int
	Name    string
	Alias   []interface{}
	Address string
	Number  []string
	Type    string
	Urls    []string
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	res := get(url + "/1")
	var post Post

	err := json.Unmarshal(res, &post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v \n", post)

	// req := Post{
	// 	UserId: 1,
	// 	Title:  "Belajar POST",
	// 	Body:   "lorem",
	// }
	// post(url, req)
}

func getOjk(id string) {
	url := "https://ojk-invest-api.vercel.com/api/illegals/" + id
	res := get(url)
	if err != nil {
		panic(err.Error())
	}
}

func get(url string) []byte {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	body := do(&client, url, "GET", nil)
	defer body.Close()

	/*
		`
			// base_url/v1/products/1
			{
				"name" : "Baju",
				"stock" : 10,
				"id" : 1
			}

			// base_url/v1/posts/1
			{
				"userId" : 1,
				"title" : "Ini title",
				"id" : 1,
				"body" : "Ini body"
			}
		`
	*/

	var post map[string]interface{}

	err := json.NewDecoder(body).Decode(&post)
	if err != nil {
		panic(err.Error())
	}

	byteData, err := json.Marshal(post)
	if err != nil {
		panic(err.Error())
	}

	return byteData
}

func post(url string, req Post) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	byteData, err := json.Marshal(req)
	if err != nil {
		panic(err.Error())
	}

	body := do(&client, url, "POST", byteData)
	defer body.Close()

	var post Post

	err = json.NewDecoder(body).Decode(post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", post)
}

func do(client *http.Client, url, method string, body []byte) io.ReadCloser {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	return res.Body
}
