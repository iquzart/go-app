package main

import (
	"os"

	"github.com/iquzart/go-app/routers"
)

func main() {

	port := ":" + os.Getenv("PORT")

	routersInit := routers.InitRouter()

	if port == ":" {
		routersInit.Run()
	} else {
		routersInit.Run(port)
	}
}
