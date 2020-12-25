package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {

	router := getRouter()
	router.HTMLRender = CreateTmplRender()
	router.GET("/", Api)
	w := performRequest(router, "GET", "/")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body is what we expect.
	//expected := `{"message":"OK"}`
	//assert.Equal(t, expected, w.Body.String())

}

func TestAbout(t *testing.T) {

	router := getRouter()
	router.HTMLRender = CreateTmplRender()
	router.GET("/about", About)
	w := performRequest(router, "GET", "/about")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body is what we expect.
	//expected := `{"message":"OK"}`
	//assert.Equal(t, expected, w.Body.String())

}

func TestUrlParam(t *testing.T) {

	router := getRouter()
	router.HTMLRender = CreateTmplRender()
	router.GET("/user/:name", UrlParam)
	w := performRequest(router, "GET", "/user/test")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestNotFound(t *testing.T) {

	router := getRouter()
	router.HTMLRender = CreateTmplRender()
	router.NoRoute(NoFound)

	w := performRequest(router, "GET", "/notfound")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusNotFound, w.Code)

}

func TestHealth(t *testing.T) {

	router := getRouter()
	router.GET("/health", Health)

	w := performRequest(router, "GET", "/health")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

}
