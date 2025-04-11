package model

// Item representa um item do pedido
type ItemDTO struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// OrderRequest representa o corpo da requisição
type OrderRequestDTO struct {
	Items []ItemDTO `json:"items"`
}

// Order representa a resposta contendo os itens com CDs alocados
type OrderDTO struct {
	Items []OrderItemDTO `json:"items"`
}

// OrderResponse representa a resposta do endpoint de alocação
type OrderResponseDTO struct {
	OrderID string   `json:"order_id,omitempty"`
	Order   OrderDTO `json:"order"`
}

// OrderItem representa um item com a informação do centro de distribuição
type OrderItemDTO struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	DistributionCenter string  `json:"distribution_center,omitempty"`
}

func ConvertToOrderDocument(response OrderResponseDTO) OrderDocument {
	var items []OrderItemDocument

	for _, dtoItem := range response.Order.Items {
		items = append(items, OrderItemDocument{
			ID:                 int64(dtoItem.ID),
			Name:               dtoItem.Name,
			Price:              dtoItem.Price,
			DistributionCenter: dtoItem.DistributionCenter,
		})
	}

	return OrderDocument{Items: items}
}
