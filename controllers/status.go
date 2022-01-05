package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http status 403 Forbidden
func Forbidden(ctx *gin.Context) {
	ctx.String(
		http.StatusForbidden,
		"Forbidden")
}

// http status 404 NotFound
func NotFound(ctx *gin.Context) {
	ctx.String(
		http.StatusNotFound,
		"Not Found")
}

// http status 405 MethodNotAllowed
func MethodNotAllowed(ctx *gin.Context) {
	ctx.String(
		http.StatusMethodNotAllowed,
		"Method Not Allowed")
}

// http status 500 InternalServerError
func InternalServerError(ctx *gin.Context) {
	ctx.String(
		http.StatusInternalServerError,
		"Application error")
}

// http status 501 NotImplemented
func NotImplemented(ctx *gin.Context) {
	ctx.String(
		http.StatusNotImplemented,
		"Not Implemented")
}

// http status 502 BadGateway
func BadGateway(ctx *gin.Context) {
	ctx.String(
		http.StatusBadGateway,
		"Bad Gateway")
}

// http status 503 ServiceUnavailable
func ServiceUnavailable(ctx *gin.Context) {
	ctx.String(
		http.StatusServiceUnavailable,
		"Application error")
}

// http status 504 GatewayTimeout
func GatewayTimeout(ctx *gin.Context) {
	ctx.String(
		http.StatusGatewayTimeout,
		"Service Unavailable")
}

// http status 505 HTTPVersionNotSupported
func HTTPVersionNotSupported(ctx *gin.Context) {
	ctx.String(
		http.StatusHTTPVersionNotSupported,
		"HTTP Version Not Supported")
}
