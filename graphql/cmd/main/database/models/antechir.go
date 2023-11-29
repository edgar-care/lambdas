package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnteChir struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name,omitempty"`
	Localisation    string             `bson:"localisation,omitempty"`
	InducedSymptoms []Symptom          `bson:"induced_symptoms,omitempty"`
}

type AnteChirCreateInput struct {
	Name            string    `bson:"name,omitempty"`
	Localisation    string    `bson:"localisation,omitempty"`
	InducedSymptoms []Symptom `bson:"induced_symptoms,omitempty"`
}

type AnteChirUpdateInput struct {
	ID              string     `bson:"_id"`
	Name            *string    `bson:"name,omitempty"`
	Localisation    *string    `bson:"localisation,omitempty"`
	InducedSymptoms *[]Symptom `bson:"induced_symptoms,omitempty"`
}
