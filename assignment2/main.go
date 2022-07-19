package main

import (
	"assignment2/routers"
)

func main() {
	PORT := ":4444"
	routers.StartServer().Run(PORT)
}
