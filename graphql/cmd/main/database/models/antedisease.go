package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnteDisease struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name,omitempty"`
	Chronicity    float64            `bson:"chronicity,omitempty"`
	SurgeryIds    *[]string          `bson:"surgery_ids,omitempty"`
	Symptoms      *[]string          `bson:"symptoms,omitempty"`
	TreatmentIds  *[]string          `bson:"treatment_ids,omitempty"`
	StillRelevant bool               `bson:"still_relevant,omitempty"`
}

type AnteDiseaseCreateInput struct {
	Name          string    `bson:"name,omitempty"`
	Chronicity    float64   `bson:"chronicity,omitempty"`
	SurgeryIds    *[]string `bson:"surgery_ids,omitempty"`
	Symptoms      *[]string `bson:"symptoms,omitempty"`
	TreatmentIds  *[]string `bson:"treatment_ids,omitempty"`
	StillRelevant bool      `bson:"still_relevant,omitempty"`
}

type AnteDiseaseUpdateInput struct {
	ID            string    `bson:"_id"`
	Name          *string   `bson:"name,omitempty"`
	Chronicity    *float64  `bson:"chronicity,omitempty"`
	SurgeryIds    *[]string `bson:"surgery_ids,omitempty"`
	Symptoms      *[]string `bson:"symptoms,omitempty"`
	TreatmentIds  *[]string `bson:"treatment_ids,omitempty"`
	StillRelevant *bool     `bson:"still_relevant,omitempty"`
}
