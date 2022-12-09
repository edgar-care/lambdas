package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     *string            `bson:"name,omitempty"`
	Age      *int32             `bson:"age,omitempty"`
}

type UserCreateInput struct {
	Email    string  `bson:"email,omitempty"`
	Password string  `bson:"password,omitempty"`
	Name     *string `bson:"name,omitempty"`
	Age      *int32  `bson:"age,omitempty"`
}

type UserUpdateInput struct {
	ID       string  `bson:"_id"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
	Name     *string `bson:"name,omitempty"`
	Age      *int32  `bson:"age,omitempty"`
}
