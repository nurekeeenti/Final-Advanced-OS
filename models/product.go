package models

type Product struct {
    Name     string `json:"name" bson:"name"`
    Price    int    `json:"price" bson:"price"`
    Quantity int    `json:"quantity" bson:"quantity"` // Добавлено поле quantity
    Category string `json:"category" bson:"category"`
    Brand    string `json:"brand" bson:"brand"`
}