package routes

import (
	"go-app/internal/interfaces/api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// addAPIRoutes adds the routes for the API controller to the specified router.
func AddSystemRoutes(r *gin.Engine) {
	// Create a new group of routes for the API controller under the specified router.
	system := r.Group("/system")

	// Add a GET route for the Health method of the System controller.
	system.GET("/health", controllers.Health)

	// Add a GET route for the Metrics endpoint using the Prometheus HTTP handler.
	system.GET("/metrics", gin.WrapH(promhttp.Handler()))

}
