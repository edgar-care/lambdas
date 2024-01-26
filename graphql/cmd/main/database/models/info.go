package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Info struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name     	string           `bson:"name,omitempty"`
	BirthDate	string              `bson:"birthdate,omitempty"`
	Height       int32              `bson:"height,omitempty"`
	Weight       int32              `bson:"weight,omitempty"`
	Sex          string             `bson:"sex,omitempty"`
	Surname string             `bson:"surname,omitempty"`
}

type InfoCreateInput struct {
	Name     	string `bson:"name,omitempty"`
	BirthDate   string    `bson:"birthdate,omitempty"`
	Height       int32    `bson:"height,omitempty"`
	Weight       int32    `bson:"weight,omitempty"`
	Sex          string   `bson:"sex,omitempty"`
	Surname string   `bson:"surname,omitempty"`
}

type InfoUpdateInput struct {
	ID           string    `bson:"_id"`
	Name     *string `bson:"name,omitempty"`
	BirthDate    *string    `bson:"birthdate,omitempty"`
	Height       *int32    `bson:"height,omitempty"`
	Weight       *int32    `bson:"weight,omitempty"`
	Sex         *string   `bson:"sex,omitempty"`
	Surname *string   `bson:"surname,omitempty"`
}