package controllers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForbidden(t *testing.T) {

	router := getRouter()
	router.GET("/403", Forbidden)

	w := performRequest(router, "GET", "/403")

	// Check the response status code is what we expect.
	assert.Equal(t, http.StatusForbidden, w.Code)

	// Check the response body is what we expect.
	expected := "Forbidden"
	assert.Equal(t, expected, w.Body.String())

}