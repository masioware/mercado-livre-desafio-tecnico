package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	group := router.Group("/api/v1")
	RegisterOrderRoutes(group)
	return router
}

func TestHandleOrder_InvalidJSON(t *testing.T) {
	router := setupTestRouter()

	reqBody := `{"items": [invalid]}`
	req, _ := http.NewRequest("POST", "/api/v1/order", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestHandleOrder_EmptyItems(t *testing.T) {
	router := setupTestRouter()

	payload := model.OrderRequestDTO{Items: []model.ItemDTO{}}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "between 1 and 100")
}

func TestHandleOrder_TooManyItems(t *testing.T) {
	router := setupTestRouter()

	items := make([]model.ItemDTO, 101)
	for i := range items {
		items[i] = model.ItemDTO{ID: i + 1, Name: "Item", Price: 1.0}
	}
	payload := model.OrderRequestDTO{Items: items}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/order", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
