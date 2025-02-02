package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Code         string             `bson:"code,omitempty"`
	Collection   string             `bson:"collection,omitempty"`
	BarCode      string             `bson:"barCode,omitempty"`
	Year         string             `bson:"year,omitempty"`
	Manufacturer string             `bson:"manufacturer,omitempty"`
	Description  string             `bson:"description,omitempty"`
	Category     string             `bson:"category,omitempty"`
}
