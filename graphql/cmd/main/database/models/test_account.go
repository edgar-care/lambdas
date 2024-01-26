package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestAccount struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type TestAccountCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type TestAccountUpdateInput struct {
	ID       string  `bson:"_id"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
}
