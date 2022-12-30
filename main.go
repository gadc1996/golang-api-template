package main

import router "gadc1996/router"

func main() {
	router := router.SetupRouter()
	router.Run(":5000")
}
