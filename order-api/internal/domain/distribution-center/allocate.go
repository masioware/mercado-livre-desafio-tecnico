package distributioncenter

import (
	"strconv"

	set "github.com/deckarep/golang-set/v2"
	model "github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
)

type DistributionCenter struct {
	Name     string
	Products set.Set[string]
}

type Catalog struct {
	ProductDistributionCenters map[int][]string
	DistributionCenterMap      map[string]*DistributionCenter
}

func NewCatalog(productDistributionCenters map[int][]string) *Catalog {
	distributionCenterMap := make(map[string]*DistributionCenter)

	for productID, dcNames := range productDistributionCenters {
		productKey := strconv.Itoa(productID)
		for _, dcName := range dcNames {
			if _, exists := distributionCenterMap[dcName]; !exists {
				distributionCenterMap[dcName] = &DistributionCenter{
					Name:     dcName,
					Products: set.NewSet[string](),
				}
			}
			distributionCenterMap[dcName].Products.Add(productKey)
		}
	}

	return &Catalog{
		ProductDistributionCenters: productDistributionCenters,
		DistributionCenterMap:      distributionCenterMap,
	}
}

func (c *Catalog) calculateCoverage(dc *DistributionCenter, unassigned set.Set[string]) int {
	count := 0
	for product := range dc.Products.Iter() {
		if unassigned.Contains(product) {
			count++
		}
	}
	return count
}

func (c *Catalog) Allocate(order model.OrderRequestDTO) model.OrderResponseDTO {
	assignments := make(map[string]string)
	unassigned := set.NewSet[string]()

	// Clona os produtos a serem alocados
	for productID := range c.ProductDistributionCenters {
		productKey := strconv.Itoa(productID)

		unassigned.Add(productKey)
	}

	// Clona o mapa de CDs para não mutar o catálogo original
	tempDCMap := make(map[string]*DistributionCenter)
	for k, v := range c.DistributionCenterMap {
		tempDCMap[k] = v
	}

	// Alocação por cobertura máxima
	for unassigned.Cardinality() > 0 && len(tempDCMap) > 0 {
		var bestDC *DistributionCenter
		maxCoverage := 0

		for _, dc := range tempDCMap {
			coverage := c.calculateCoverage(dc, unassigned)
			if coverage > maxCoverage {
				bestDC = dc
				maxCoverage = coverage
			}
		}

		if bestDC == nil {
			break
		}

		for product := range bestDC.Products.Iter() {
			if unassigned.Contains(product) {
				assignments[product] = bestDC.Name
				unassigned.Remove(product)
			}
		}

		delete(tempDCMap, bestDC.Name)
	}

	// Monta o OrderResponse
	var response model.OrderResponseDTO
	for _, item := range order.Items {
		productKey := strconv.Itoa(item.ID)
		dc := assignments[productKey]
		response.Order.Items = append(response.Order.Items, model.OrderItemDTO{
			ID:                 item.ID,
			Name:               item.Name,
			Price:              item.Price,
			DistributionCenter: dc,
		})
	}

	return response
}
