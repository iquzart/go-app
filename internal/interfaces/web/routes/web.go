package routes

import (
	"go-app/internal/interfaces/web/controllers"
	"go-app/internal/interfaces/web/helpers"

	"github.com/gin-gonic/gin"
)

// AddWebRoutes adds the routes for the web controller to the specified router.
func AddWebRoutes(r *gin.Engine) {

	// Serve static files under "/static" from the "./public/static" directory.
	r.Static("/static", "./public/static")

	// Set up the HTML template rendering.
	r.HTMLRender = helpers.CreateTmplRender()

	// Create a new group of routes for the web controller under the specified router.
	web := r.Group("/")

	// Add a GET route for the Home method of the web controller.
	web.GET("/", controllers.Home)
	// Add a GET route for the About method of the web controller.
	web.GET("/about", controllers.About)
}
