package controllers

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	banner := os.Getenv("BANNER")
	hostname, _ := os.Hostname()
	ip, _ := net.LookupIP(hostname)
	ctx.HTML(
		http.StatusOK,
		"home",
		gin.H{
			"title":    "Home",
			"hostname": hostname,
			"IP":       ip,
			"banner":   banner,
		},
	)

}

func About(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"about",
		gin.H{
			"title": "About",
		},
	)

}

func Health(ctx *gin.Context) {
	ctx.String(
		http.StatusOK,
		"Working!")
}

func UrlParam(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.HTML(
		http.StatusOK,
		"name",
		gin.H{
			"title": "user",
			"name":  name,
		},
	)

}

func NoFound(ctx *gin.Context) {
	ctx.HTML(
		404,
		"404",
		gin.H{
			"title": "404",
		},
	)
}
