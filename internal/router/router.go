package router

import (
	apiRoutes "go-app/internal/interfaces/api/routes"
	webRoutes "go-app/internal/interfaces/web/routes"
	"go-app/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and sets up the middleware and routes.
func SetupRouter(serviceName string) *gin.Engine {
	// Create a new Gin router instance.
	router := gin.New()

	// Add Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Use the Prometheus middleware to instrument the router with metrics.
	router.Use(middlewares.PrometheusMetrics(serviceName))

	// Add the routes to the router.
	addRoutes(router)

	// Return the initialized Gin router.
	return router
}

// addRoutes adds the system and application routes to the specified router.
func addRoutes(r *gin.Engine) {
	// Add the routes for the System.
	apiRoutes.AddSystemRoutes(r)
	// Add the routes for the API.
	apiRoutes.AddAPIRoutes(r)
	// Add HTTP Status
	apiRoutes.AddStatusRoutes(r)
	// add Web Routes
	webRoutes.AddWebRoutes(r)
}
