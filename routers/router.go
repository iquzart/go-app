package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/iquzart/go-app/controllers"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
)

func CreateTmplRender() multitemplate.Renderer {

	baseTmpl := "public/templates/base.tmpl.html"
	homeTmpl := "public/templates/home.tmpl.html"
	aboutTmpl := "public/templates/about.tmpl.html"
	nameTmpl := "public/templates/name.tmpl.html"
	s404Tmpl := "public/templates/404.tmpl.html"

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
	p := ginprometheus.NewPrometheus("gin")

	p.Use(r)

	// Frontend
	r.Static("/static", "./public/static")
	r.HTMLRender = CreateTmplRender()

	// Routes
	r.GET("/", controllers.Home)
	r.GET("/about", controllers.About)
	r.GET("/api", controllers.Api)
	r.GET("/user/:name", controllers.UrlParam)
	r.GET("/health", controllers.Health)

	r.GET("/403", controllers.Forbidden)	
	r.GET("/404", controllers.NotFound)	
	r.GET("/405", controllers.MethodNotAllowed)		
	r.GET("/500", controllers.InternalServerError)
	r.GET("/501", controllers.NotImplemented)
	r.GET("/502", controllers.BadGateway)
	r.GET("/503", controllers.ServiceUnavailable)
	r.GET("/504", controllers.GatewayTimeout)
	r.GET("/505", controllers.HTTPVersionNotSupported)

	r.NoRoute(controllers.NoFound)

	return r
}
