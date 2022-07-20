package main

import "jwt/router"

func main() {
	r := router.StartApp()

	r.Run(":4444")
}
