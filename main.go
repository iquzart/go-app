package main

import (
	"github.com/iquzart/go-app/router"
)

func main() {

	routersInit := router.InitRouter()

	routersInit.Run(":8080")

}
