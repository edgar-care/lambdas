package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Logs struct {
	Question string `bson:"question,omitempty"`
	Answer   string `bson:"answer,omitempty"`
}

type Session struct {
	ID           primitive.ObjectID `bson:"_id"`
	Symptoms     []string           `bson:"symptoms,omitempty"`
	Age          int32              `bson:"age,omitempty"`
	Height       int32              `bson:"height,omitempty"`
	Weight       int32              `bson:"weight,omitempty"`
	Sex          string             `bson:"sex,omitempty"`
	LastQuestion string             `bson:"last_question,omitempty"`
	Logs         []Logs             `bson:"logs,omitempty"`
}

type SessionCreateInput struct {
	Symptoms     []string `bson:"symptoms,omitempty"`
	Age          int32    `bson:"age,omitempty"`
	Height       int32    `bson:"height,omitempty"`
	Weight       int32    `bson:"weight,omitempty"`
	Sex          string   `bson:"sex,omitempty"`
	LastQuestion string   `bson:"last_question,omitempty"`
	Logs         []Logs   `bson:"logs,omitempty"`
}

type SessionUpdateInput struct {
	ID           string    `bson:"_id"`
	Symptoms     *[]string `bson:"symptoms,omitempty"`
	Age          *int32    `bson:"age,omitempty"`
	Height       *int32    `bson:"height,omitempty"`
	Weight       *int32    `bson:"weight,omitempty"`
	Sex          *string   `bson:"sex,omitempty"`
	LastQuestion *string   `bson:"last_question,omitempty"`
	Logs         *[]Logs   `bson:"logs,omitempty"`
}
