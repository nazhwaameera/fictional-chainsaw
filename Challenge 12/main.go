package main

import (
	"Challenge_12/database"
	"Challenge_12/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
