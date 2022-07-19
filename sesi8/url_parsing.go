package main

import (
	"fmt"
	"net/url"
)

func main() {
	uri := "https://www.noobee.id/seacrh?q=golang&level=basic&level=medium"

	url, err := url.Parse(uri)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("path:", url.Path)
	fmt.Println("rawPath:", url.RawPath)
	fmt.Println("schema:", url.Scheme)
	fmt.Println("host:", url.Host)
	fmt.Println("query:", url.Query())
	fmt.Println("rawQuery:", url.RawQuery)
}
