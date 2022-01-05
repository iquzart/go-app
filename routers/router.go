package routers

import (
	helpers "go-app/helpers"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	// Creates a router without any middleware
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
	initializeRoutes(r)

	return r
}
