package main

import "final-project/routers"

func main() {
	PORT := ":4444"
	routers.StartServer().Run(PORT)
}
