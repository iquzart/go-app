package main

import (
	"os"
	"log"
	"go-app/routers"
)

func main() {
	// Custom port 
	port := ":" + os.Getenv("PORT")

	// Initialize the router
	routersInit := routers.InitRouter()

	// Run Application
	if port == ":" {
		log.Printf("Starting app on default port")
		routersInit.Run()
	} else {
		log.Printf("Starting app on port " + os.Getenv("PORT"))
		routersInit.Run(port)
	}


}
