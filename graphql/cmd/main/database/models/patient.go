package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	LastName string             `bson:"last_name,omitempty"`
	Age      int32              `bson:"age,omitempty"`
	Height   int32              `bson:"height,omitempty"`
	Weight   int32              `bson:"weight,omitempty"`
	Sex      string             `bson:"sex,omitempty"`
}

type PatientCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Name     string `bson:"name,omitempty"`
	LastName string `bson:"last_name,omitempty"`
	Age      int32  `bson:"age,omitempty"`
	Height   int32  `bson:"height,omitempty"`
	Weight   int32  `bson:"weight,omitempty"`
	Sex      string `bson:"sex,omitempty"`
}

type PatientUpdateInput struct {
	ID       string  `bson:"_id"`
	Email    *string `bson:"email,omitempty"`
	Password *string `bson:"password,omitempty"`
	Name     *string `bson:"name,omitempty"`
	LastName *string `bson:"last_name,omitempty"`
	Age      *int32  `bson:"age,omitempty"`
	Height   *int32  `bson:"height,omitempty"`
	Weight   *int32  `bson:"weight,omitempty"`
	Sex      *string `bson:"sex,omitempty"`
}
