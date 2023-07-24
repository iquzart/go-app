package routes

import (
	"go-app/internal/interfaces/api/controllers"

	"github.com/gin-gonic/gin"
)

// addAPIRoutes adds the routes for the API controller to the specified router.
func AddAPIRoutes(r *gin.Engine) {
	// Create a new group of routes for the API controller under the specified router.
	api := r.Group("/api")

	// Add a GET route for the APIVersion method of the API controller.
	api.GET("/version", controllers.APIVersion)
}
