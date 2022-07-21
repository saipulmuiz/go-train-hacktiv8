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

type ResKBBI struct {
	Status bool
	Data   QueryKBBI
}

type QueryKBBI struct {
	Total       int
	ResultLists []KBBI
}

type KBBI struct {
	Kata string
	Arti []string
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
	// getKBBI("apel")
	getOJK("1")
	// getPost("2")

	// req := Post{
	// 	UserId: 1,
	// 	Title:  "Belajar POST",
	// 	Body:   "lorem",
	// }
	// post(url, req)
}

func getKBBI(search string) {
	url := "https://kbbi-api-amm.herokuapp.com/search?q=" + search
	res := get(url)

	var resKbbi ResKBBI
	err := json.Unmarshal(res, &resKbbi)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v \n", resKbbi)
}

func getPost(id string) {
	url := "https://jsonplaceholder.typicode.com/posts/" + id
	res := get(url)
	var post Post

	err := json.Unmarshal(res, &post)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v \n", post)
}

func getOJK(id string) {
	url := "https://ojk-invest-api.vercel.app/api/illegals/" + id
	res := get(url)

	var ojk ResOJKSingle
	err := json.Unmarshal(res, &ojk)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v \n", ojk)
}

func get(url string) []byte {
	client := http.Client{
		// Timeout: 10000 * time.Millisecond,
	}
	body := do(&client, url, "GET", nil)
	defer body.Close()
	var result map[string]interface{}

	err := json.NewDecoder(body).Decode(&result)
	if err != nil {
		panic(err.Error())
	}

	byteData, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}
	return byteData
}

func post(url string, req Post) {
	client := http.Client{
		Timeout: 2000 * time.Millisecond,
	}

	byteData, err := json.Marshal(req)
	if err != nil {
		panic(err.Error())
	}

	body := do(&client, url, "POST", byteData)
	defer body.Close()

	var post Post

	err = json.NewDecoder(body).Decode(&post)
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
