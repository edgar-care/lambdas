package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Disease struct {
	ID               primitive.ObjectID `bson:"_id"`
	Code             string             `bson:"code,omitempty"`
	Name             string             `bson:"name,omitempty"`
	Symptoms         []string           `bson:"symptoms,omitempty"`
	SymptomsAcute    *[]SymptomWeight   `bson:"symptoms_acute,omitempty"`
	SymptomsSubacute *[]SymptomWeight   `bson:"symptoms_subacute,omitempty"`
	SymptomsChronic  *[]SymptomWeight   `bson:"symptoms_chronic,omitempty"`
	Advice           *string            `bson:"advice,omitempty"`
}

type DiseaseCreateInput struct {
	Code             string           `bson:"code,omitempty"`
	Name             string           `bson:"name,omitempty"`
	Symptoms         []string         `bson:"symptoms,omitempty"`
	SymptomsAcute    *[]SymptomWeight `bson:"symptoms_acute,omitempty"`
	SymptomsSubacute *[]SymptomWeight `bson:"symptoms_subacute,omitempty"`
	SymptomsChronic  *[]SymptomWeight `bson:"symptoms_chronic,omitempty"`
	Advice           *string          `bson:"advice,omitempty"`
}

type DiseaseUpdateInput struct {
	ID               string           `bson:"_id"`
	Code             *string          `bson:"code,omitempty"`
	Name             *string          `bson:"name,omitempty"`
	Symptoms         *[]string        `bson:"symptoms,omitempty"`
	SymptomsAcute    *[]SymptomWeight `bson:"symptoms_acute,omitempty"`
	SymptomsSubacute *[]SymptomWeight `bson:"symptoms_subacute,omitempty"`
	SymptomsChronic  *[]SymptomWeight `bson:"symptoms_chronic,omitempty"`
	Advice           *string          `bson:"advice,omitempty"`
}
