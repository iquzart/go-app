package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/iquzart/go-app/controllers"
)

func CreateTmplRender() multitemplate.Renderer {

	baseTmpl := "frontend/templates/base.tmpl.html"
	homeTmpl := "frontend/templates/home.tmpl.html"
	aboutTmpl := "frontend/templates/about.tmpl.html"
	nameTmpl := "frontend/templates/name.tmpl.html"
	s404Tmpl := "frontend/templates/404.tmpl.html"

	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", baseTmpl, homeTmpl)
	r.AddFromFiles("404", baseTmpl, s404Tmpl)
	r.AddFromFiles("about", baseTmpl, aboutTmpl)
	r.AddFromFiles("name", baseTmpl, nameTmpl)
	return r
}

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Frontend
	r.Static("/static", "./frontend/static")
	r.HTMLRender = CreateTmplRender()

	// Routes
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)
	r.GET("/api", controllers.Api)
	r.GET("/user/:name", controllers.ApiParam)
	r.GET("/health", controllers.Health)
	r.NoRoute(controllers.NoFound)

	return r
}
