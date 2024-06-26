package main

import (
	"fmt"
	"gin/database"
	router "gin/router"
)

func main() {
	err := database.ConnectDatabase()
	if err == nil {
		fmt.Println("\n Database connected!")
	} else {
		fmt.Println("Fatal error database connection", err)
	}

	r := router.SetupRouter()
	r.Run(":8088")
}
