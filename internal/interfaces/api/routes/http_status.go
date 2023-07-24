package routes

import (
	"go-app/internal/interfaces/api/controllers"

	"github.com/gin-gonic/gin"
)

// AddStatusRoutes adds the routes for the API controller to the specified router.
func AddStatusRoutes(r *gin.Engine) {
	// Create a new group of routes for the API controller under the specified router.
	status := r.Group("/api/status")

	// Add routes for the HTTP Status methods.
	status.GET("/403", controllers.Forbidden)
	status.GET("/404", controllers.NotFound)
	status.GET("/405", controllers.MethodNotAllowed)
	status.GET("/500", controllers.InternalServerError)
	status.GET("/501", controllers.NotImplemented)
	status.GET("/502", controllers.BadGateway)
	status.GET("/503", controllers.ServiceUnavailable)
	status.GET("/504", controllers.GatewayTimeout)
	status.GET("/505", controllers.HTTPVersionNotSupported)
}
