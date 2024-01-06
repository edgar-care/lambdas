package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnteDisease struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name,omitempty"`
	Chronicity float64            `bson:"chronicity,omitempty"`
	Chir       *string            `bson:"chir,omitempty"`
	Treatment  *[]string          `bson:"treatment,omitempty"`
	Symptoms   *[]string          `bson:"symptoms,omitempty"`
}

type AnteDiseaseCreateInput struct {
	Name       string    `bson:"name,omitempty"`
	Chronicity float64   `bson:"chronicity,omitempty"`
	Chir       *string   `bson:"chir,omitempty"`
	Treatment  *[]string `bson:"treatment,omitempty"`
	Symptoms   *[]string `bson:"symptoms,omitempty"`
}

type AnteDiseaseUpdateInput struct {
	ID         string    `bson:"_id"`
	Name       *string   `bson:"name,omitempty"`
	Chronicity *float64  `bson:"chronicity,omitempty"`
	Chir       *string   `bson:"chir,omitempty"`
	Treatment  *[]string `bson:"treatment,omitempty"`
	Symptoms   *[]string `bson:"symptoms,omitempty"`
}
