package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	DOB         string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	CreatedAt   string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
