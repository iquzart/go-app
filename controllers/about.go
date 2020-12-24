package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func About(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"about",
		gin.H{
			"title": "About",
		},
	)

}
