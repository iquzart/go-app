package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iquzart/go-app/routers"
	"github.com/stretchr/testify/assert"
)

func getRouter() *gin.Engine {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	r := routers.InitRouter()
	return r
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealth(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/health")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body is what we expect.
	expected := `Working!`
	assert.Contains(t, expected, w.Body.String())

}

func TestApi(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/api")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body is what we expect.
	expected := `{"message":"OK"}`
	assert.Equal(t, expected, w.Body.String())

}
func TestApiPram(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/user/test")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAbout(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/about")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestHome(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestNotFound(t *testing.T) {

	router := getRouter()

	w := performRequest(router, "GET", "/notfound")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusNotFound, w.Code)

}
