package helpers

import (
	"github.com/gin-contrib/multitemplate"
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
