package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/repository"

	domain "github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/domain/distribution-center"
	model "github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
)

func RegisterOrderRoutes(rg *gin.RouterGroup) {
	rg.POST("/order", HandleProcessOrder)
	rg.GET("/order/:id", HandleGetOrderByID)
}

// HandleOrder godoc
// @Summary Processa um pedido e aloca centros de distribuição
// @Description Recebe uma lista de itens e retorna os centros de distribuição que irão atendê-los
// @Tags order
// @Accept json
// @Produce json
// @Param order body model.OrderRequest true "Dados do pedido"
// @Success 200 {object} model.OrderResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /order [post]
func HandleProcessOrder(c *gin.Context) {
	var order model.OrderRequestDTO

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(order.Items) == 0 || len(order.Items) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "order must contain between 1 and 100 items",
		})
		return
	}

	distributionResults, errors := domain.RetrieveDistributionCenters(order)
	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch distribution centers",
		})
		return
	}

	catalog := domain.NewCatalog(distributionResults)
	response := catalog.Allocate(order)

	document := model.ConvertToOrderDocument(response)
	id, err := repository.SaveOrderDocument(document)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save order request",
			"error":   err.Error(),
		})
		return
	}

	response.OrderID = id.Hex()

	c.JSON(http.StatusOK, response)
}

// HandleGetOrderByID godoc
// @Summary Busca um pedido salvo pelo ID
// @Description Retorna um pedido previamente salvo no MongoDB
// @Tags order
// @Produce json
// @Param id path string true "ID do pedido"
// @Success 200 {object} model.OrderResponseWithID
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /order/{id} [get]
func HandleGetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := repository.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
