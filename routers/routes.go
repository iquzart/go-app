package routers

import (
	controllers "go-app/controllers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	// Routes
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)
	r.GET("/user/:name", controllers.UrlParam)
	r.GET("/health", controllers.Health)
	r.NoRoute(controllers.NoFound)

	// Route Group - Status
	status := r.Group("/status")
	{
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

	// Route Group - API
	v1 := r.Group("/api/v1")
	{
		v1.GET("/check", controllers.Api)
	}
}
