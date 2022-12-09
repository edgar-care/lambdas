package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"last_name,omitempty"`
}

type DoctorCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Name     string `bson:"name,omitempty"`
	LastName string `bson:"last_name,omitempty"`
}

type DoctorUpdateInput struct {
	ID       string  `bson:"_id"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
	Name     *string `bson:"name,omitempty"`
	LastName *string `bson:"last_name,omitempty"`
}
