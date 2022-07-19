package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	var post []Post
	// byteRes, err := ioutil.ReadAll(res.Body)
	err = json.NewDecoder(res.Body).Decode(&post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", post[1])
}
