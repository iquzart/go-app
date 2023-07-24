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
