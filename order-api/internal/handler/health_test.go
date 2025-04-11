package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	// Configura o router isoladamente
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	group := router.Group("/api/v1")
	RegisterHealthCheckerRoutes(group)

	// Cria a requisição simulada
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/health-check", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "ok")
	assert.Contains(t, resp.Body.String(), "Service is healthy")
}
