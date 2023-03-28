package main

import (
	"Challenge_8/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
