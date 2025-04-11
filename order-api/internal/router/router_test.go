package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter_HealthCheckRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := NewRouter()

	req, _ := http.NewRequest("GET", "/api/v1/health-check", nil)
	resp := httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "healthy")
}

func TestNewRouter_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app := NewRouter()

	req, _ := http.NewRequest("GET", "/api/v1/nonexistent", nil)
	resp := httptest.NewRecorder()

	app.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
