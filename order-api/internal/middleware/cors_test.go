package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSMiddleware_OptionsMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORSMiddleware())
	router.OPTIONS("/test", func(c *gin.Context) {})

	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 204, resp.Code)
	assert.Equal(t, "*", resp.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, resp.Header().Get("Access-Control-Allow-Methods"), "OPTIONS")
	assert.Contains(t, resp.Header().Get("Access-Control-Allow-Headers"), "Authorization")
}

func TestCORSMiddleware_NonOptionsRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CORSMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	req, _ := http.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "*", resp.Header().Get("Access-Control-Allow-Origin"))
	assert.Contains(t, resp.Header().Get("Access-Control-Allow-Methods"), "GET")
}
