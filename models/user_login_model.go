package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserLoginModel struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
