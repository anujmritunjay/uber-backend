package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SignUp struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" validate:"required"`
	LastName  string             `json:"lastName" validate:"required"`
	Email     string             `json:"email" validate:"required,email"`
	Password  string             `json:"password" validate:"required,min=6"`
}
