package repository

import (
	"context"
	"errors"
	"time"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/config"
	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveOrderDocument(order model.OrderDocument) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := config.MongoDB.Collection("orders")
	res, err := coll.InsertOne(ctx, order)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func GetOrderByID(id string) (model.OrderResponseDTO, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.OrderResponseDTO{}, errors.New("invalid order ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var doc model.OrderDocument

	err = config.MongoDB.Collection("orders").FindOne(ctx, bson.M{"_id": objectID}).Decode(&doc)
	if err != nil {
		return model.OrderResponseDTO{}, err
	}

	// Converter para DTO de resposta
	var items []model.OrderItemDTO
	for _, item := range doc.Items {
		items = append(items, model.OrderItemDTO{
			ID:                 int(item.ID),
			Name:               item.Name,
			Price:              item.Price,
			DistributionCenter: item.DistributionCenter,
		})
	}

	return model.OrderResponseDTO{
		OrderID: objectID.Hex(),
		Order: model.OrderDTO{
			Items: items,
		},
	}, nil
}
