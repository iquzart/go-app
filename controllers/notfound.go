package controllers

import (
	"github.com/gin-gonic/gin"
)

func NoFound(ctx *gin.Context) {
	ctx.HTML(
		404,
		"404",
		gin.H{
			"title": "404",
		},
	)
}
