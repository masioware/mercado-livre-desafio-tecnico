package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDocument struct {
	ID    primitive.ObjectID  `bson:"_id,omitempty"`
	Items []OrderItemDocument `bson:"items"`
}

type OrderItemDocument struct {
	ID                 int64   `bson:"id"`
	Name               string  `bson:"name"`
	Price              float64 `bson:"price"`
	DistributionCenter string  `bson:"distribution_center"`
}
