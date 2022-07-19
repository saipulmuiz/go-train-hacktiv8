package main

import "belajar-gin/routers"

func main() {
	const PORT = ":8000"

	routers.StartServer().Run(PORT)
}
