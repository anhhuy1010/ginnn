package main

import (
	"fmt"
	"students/database"
	"students/router"
)

func main() {
	err := database.ConnectDatabase()
	if err == nil {
		fmt.Println("\n Database connected!")
	} else {
		fmt.Println("Fatal error database connection", err)
	}

	r := router.SetupRouter()
	r.Run(":7299")
}
