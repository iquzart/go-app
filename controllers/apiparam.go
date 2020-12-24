package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiParam(ctx *gin.Context) {
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
