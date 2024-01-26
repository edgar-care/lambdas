package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Treatment struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name,omitempty"`
	Disease     string             `bson:"disease,omitempty"`
	Symptoms    *[]string          `bson:"symptoms,omitempty"`
	SideEffects *[]string          `bson:"side_effects,omitempty"`
}

type TreatmentCreateInput struct {
	Name        string    `bson:"name,omitempty"`
	Disease     string    `bson:"disease,omitempty"`
	Symptoms    *[]string `bson:"symptoms,omitempty"`
	SideEffects *[]string `bson:"side_effects,omitempty"`
}

type TreatmentUpdateInput struct {
	ID          string    `bson:"_id"`
	Name        *string   `bson:"name,omitempty"`
	Disease     *string   `bson:"disease,omitempty"`
	Symptoms    *[]string `bson:"symptoms,omitempty"`
	SideEffects *[]string `bson:"side_effects,omitempty"`
}
