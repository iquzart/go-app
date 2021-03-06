package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {

	router := getRouter()
	router.GET("/api", Api)

	w := performRequest(router, "GET", "/api")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body is what we expect.
	expected := `{"message":"OK"}`
	assert.Equal(t, expected, w.Body.String())

}
