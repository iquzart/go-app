package server

import (
	helpers "go-app/helpers"
	routes "go-app/routes"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

// InitServer initialize routing information
func InitServer() *gin.Engine {

	// Creates a server without any middleware
	r := gin.New()

	// Add Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)

	// Frontend
	r.Static("/static", "./public/static")
	r.HTMLRender = helpers.CreateTmplRender()

	// Initialize Routes
	routes.InitializeRoutes(r)

	return r
}
