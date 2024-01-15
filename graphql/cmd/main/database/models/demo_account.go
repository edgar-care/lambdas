package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DemoAccount struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type DemoAccountCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type DemoAccountUpdateInput struct {
	ID       string  `bson:"_id"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
}
