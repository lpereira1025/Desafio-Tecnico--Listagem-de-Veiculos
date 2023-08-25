package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHomePage(t *testing.T) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Página Inicial")
	})

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", w.Code)
	}

	expectedBody := "Página Inicial"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, w.Body.String())
	}
}
