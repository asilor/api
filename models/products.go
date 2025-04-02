package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Product struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Price       float64       `json:"price" bson:"price"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
}
