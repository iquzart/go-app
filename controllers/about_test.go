package contorllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAbout(t *testing.T) {
	router := gin.Default()
	router.Get("/about", About)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/about", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("You recieved a %v error.", w.Code)
	}

}