package main

import "gin/routers"

func main() {
	const PORT = ":4444"

	routers.StartServer().Run(PORT)
}
