package main

import (
	router "gin/router"
)

func main() {


	r := router.SetupRouter()
	r.Run(":8088")
}
