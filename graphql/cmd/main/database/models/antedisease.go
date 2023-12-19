package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnteDisease struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name,omitempty"`
	Chronicity float32            `bson:"chronicity,omitempty"`
	Chir       AnteChir           `bson:"chir,omitempty"`
	Treatment  []Treatment        `bson:"treatment,omitempty"`
	Symptoms   []Symptom          `bson:"symptoms,omitempty"`
}

type AnteDiseaseCreateInput struct {
	Name       string      `bson:"name,omitempty"`
	Chronicity float32     `bson:"chronicity,omitempty"`
	Chir       AnteChir    `bson:"chir,omitempty"`
	Treatment  []Treatment `bson:"treatment,omitempty"`
	Symptoms   []Symptom   `bson:"symptoms,omitempty"`
}

type AnteDiseaseUpdateInput struct {
	ID         string       `bson:"_id"`
	Name       *string      `bson:"name,omitempty"`
	Chronicity *float32     `bson:"chronicity,omitempty"`
	Chir       *AnteChir    `bson:"chir,omitempty"`
	Treatment  *[]Treatment `bson:"treatment,omitempty"`
	Symptoms   *[]Symptom   `bson:"symptoms,omitempty"`
}
