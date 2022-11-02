package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {

	router := getRouter()
	router.GET("/500", InternalServerError)

	w := performRequest(router, "GET", "/500")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Check the response body is what we expect.
	expected := "Application error"
	assert.Equal(t, expected, w.Body.String())

}
