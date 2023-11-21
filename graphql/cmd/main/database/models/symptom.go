package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Symptom struct {
	ID       primitive.ObjectID `bson:"_id"`
	Code     string             `bson:"code,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Symptom  []string           `bson:"symptom,omitempty"`
	Advice   *string            `bson:"advice,omitempty"`
	Question string             `bson:"question,omitempty"`
}

type SymptomCreateInput struct {
	Code     string   `bson:"code,omitempty"`
	Name     string   `bson:"name,omitempty"`
	Symptom  []string `bson:"symptom,omitempty"`
	Advice   *string  `bson:"advice,omitempty"`
	Question string   `bson:"question,omitempty"`
}

type SymptomUpdateInput struct {
	ID       string    `bson:"_id"`
	Code     *string   `bson:"code,omitempty"`
	Name     string    `bson:"name,omitempty"`
	Symptom  *[]string `bson:"symptom,omitempty"`
	Advice   *string   `bson:"advice,omitempty"`
	Question *string   `bson:"question,omitempty"`
}
