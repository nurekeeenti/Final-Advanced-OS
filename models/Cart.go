package models

import "go.mongodb.org/mongo-driver/bson/primitive"



type Cart struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Products []Product          `json:"products" bson:"products"`
}

// type Product struct {
// 	Name     string  `json:"name" bson:"name"`
// 	Price    float64 `json:"price" bson:"price"`
// 	Quantity int     `json:"quantity" bson:"quantity"`
// }