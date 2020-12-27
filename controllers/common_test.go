package controllers

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func CreateTmplRender() multitemplate.Renderer {

	baseTmpl := "../public/templates/base.tmpl.html"
	homeTmpl := "../public/templates/home.tmpl.html"
	aboutTmpl := "../public/templates/about.tmpl.html"
	nameTmpl := "../public/templates/name.tmpl.html"
	s404Tmpl := "../public/templates/404.tmpl.html"

	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", baseTmpl, homeTmpl)
	r.AddFromFiles("404", baseTmpl, s404Tmpl)
	r.AddFromFiles("about", baseTmpl, aboutTmpl)
	r.AddFromFiles("name", baseTmpl, nameTmpl)
	return r
}
