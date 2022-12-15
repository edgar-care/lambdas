package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Health struct {
	ID           primitive.ObjectID `bson:"_id"`
	Patientallergies     	string           `bson:"patientallergies,omitempty"`
	Patientsillness string             `bson:"patientsillness,omitempty"`
}

type HealthCreateInput struct {
	Patientallergies     	string `bson:"patientallergies,omitempty"`
	Patientsillness string   `bson:"surname,omitempty"`
}

type HealthUpdateInput struct {
	ID           string    `bson:"_id"`
	Patientallergies     *string `bson:"patientallergies,omitempty"`
	Patientsillness *string   `bson:"surname,omitempty"`
}