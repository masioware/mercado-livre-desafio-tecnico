package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHealthCheckerRoutes registra as rotas de verificação de saúde do serviço
func RegisterHealthCheckerRoutes(rg *gin.RouterGroup) {
	rg.GET("/health-check", HealthCheck)
}

// HealthCheck godoc
// @Summary Verifica se o serviço está funcionando
// @Description Endpoint para checagem de saúde da API
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Service is healthy",
	})
}
