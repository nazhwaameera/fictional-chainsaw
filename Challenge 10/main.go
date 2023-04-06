package main

import (
	"Challenge_10/database"
	"Challenge_10/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
