package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	data := map[string]interface{}{
		"title":  "Test doang",
		"body":   "Belajar post",
		"userId": 1,
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	client := http.Client{
		Timeout: 1000 * time.Millisecond,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	var post Post

	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", post)

}
