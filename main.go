package main

import (
	"gadc1996/database"
	router "gadc1996/router"
)

func main() {
	database := new(database.Database)
	database.Setup()
	router := router.SetupRouter()
	router.Run(":5000")
}
