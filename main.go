package main

import (
	"go-app/server"
	"log"
	"os"
)

func main() {
	// Custom port
	port := ":" + os.Getenv("PORT")

	// Initialize the router
	routersInit := server.InitServer()

	// Run Application
	if port == ":" {
		log.Printf("Go App started on default port")
		routersInit.Run()
	} else {
		log.Printf("Go App started on port " + os.Getenv("PORT"))
		routersInit.Run(port)
	}

}
