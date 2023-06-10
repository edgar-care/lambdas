package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `bson:"_id"`
	Symptoms     []string           `bson:"symptoms,omitempty"`
	Age          int32              `bson:"age,omitempty"`
	Height       int32              `bson:"height,omitempty"`
	Weight       int32              `bson:"weight,omitempty"`
	Sex          string             `bson:"sex,omitempty"`
	LastQuestion string             `bson:"last_question,omitempty"`
}

type SessionCreateInput struct {
	Symptoms     []string `bson:"symptoms,omitempty"`
	Age          int32    `bson:"age,omitempty"`
	Height       int32    `bson:"height,omitempty"`
	Weight       int32    `bson:"weight,omitempty"`
	Sex          string   `bson:"sex,omitempty"`
	LastQuestion string   `bson:"last_question,omitempty"`
}

type SessionUpdateInput struct {
	ID           string    `bson:"_id"`
	Symptoms     *[]string `bson:"symptoms,omitempty"`
	Age          *int32    `bson:"age,omitempty"`
	Height       *int32    `bson:"height,omitempty"`
	Weight       *int32    `bson:"weight,omitempty"`
	Sex          *string   `bson:"sex,omitempty"`
	LastQuestion *string   `bson:"last_question,omitempty"`
}
