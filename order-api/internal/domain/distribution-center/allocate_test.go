package distributioncenter

import (
	"testing"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCatalog(t *testing.T) {
	input := map[int][]string{
		1: {"CD1", "CD2"},
		2: {"CD2"},
	}

	catalog := NewCatalog(input)

	assert.Len(t, catalog.ProductDistributionCenters, 2)
	assert.Contains(t, catalog.DistributionCenterMap, "CD1")
	assert.Contains(t, catalog.DistributionCenterMap, "CD2")
}

func TestAllocate_AssignsBestDistributionCenters(t *testing.T) {
	productCDs := map[int][]string{
		1: {"CD1", "CD2"},
		2: {"CD2"},
		3: {"CD3"},
	}

	order := model.OrderRequestDTO{
		Items: []model.ItemDTO{
			{ID: 1, Name: "Item 1", Price: 10.0},
			{ID: 2, Name: "Item 2", Price: 20.0},
			{ID: 3, Name: "Item 3", Price: 30.0},
		},
	}

	catalog := NewCatalog(productCDs)
	response := catalog.Allocate(order)

	assert.Len(t, response.Order.Items, 3)
	for _, item := range response.Order.Items {
		assert.NotEmpty(t, item.DistributionCenter)
	}
}

func TestAllocate_WithNoAvailableCDs(t *testing.T) {
	productCDs := map[int][]string{
		1: {},
	}

	order := model.OrderRequestDTO{
		Items: []model.ItemDTO{
			{ID: 1, Name: "Item 1", Price: 10.0},
		},
	}

	catalog := NewCatalog(productCDs)
	response := catalog.Allocate(order)

	assert.Len(t, response.Order.Items, 1)
	assert.Empty(t, response.Order.Items[0].DistributionCenter)
}
