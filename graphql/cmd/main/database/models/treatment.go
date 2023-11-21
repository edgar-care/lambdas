package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Treatment struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name,omitempty"`
	Disease     Disease            `bson:"disease,omitempty"`
	Symptoms    []Symptom          `bson:"symptoms,omitempty"`
	SideEffects []Symptom          `bson:"sideeffects,omitempty"`
}

type TreatmentCreateInput struct {
	Name        string    `bson:"name,omitempty"`
	Disease     Disease   `bson:"disease,omitempty"`
	Symptoms    []Symptom `bson:"symptoms,omitempty"`
	SideEffects []Symptom `bson:"sideeffects,omitempty"`
}

type TreatmentUpdateInput struct {
	ID          string     `bson:"_id"`
	Name        *string    `bson:"name,omitempty"`
	Disease     *Disease   `bson:"disease,omitempty"`
	Symptoms    *[]Symptom `bson:"symptoms,omitempty"`
	SideEffects *[]Symptom `bson:"sideeffects,omitempty"`
}
