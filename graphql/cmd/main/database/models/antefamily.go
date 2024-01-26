package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnteFamily struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name,omitempty"`
	Disease []string           `bson:"disease,omitempty"`
}

type AnteFamilyCreateInput struct {
	Name    string   `bson:"name,omitempty"`
	Disease []string `bson:"disease,omitempty"`
}

type AnteFamilyUpdateInput struct {
	ID      string    `bson:"_id"`
	Name    *string   `bson:"name,omitempty"`
	Disease *[]string `bson:"disease,omitempty"`
}
