package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SymptomWeight struct {
	Key   string  `bson:"key"`
	Value float64 `bson:"value,omitempty"`
}

type Symptom struct {
	ID       primitive.ObjectID `bson:"_id"`
	Code     string             `bson:"code,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Location *string            `bson:"location,omitempty"`
	Duration *int32             `bson:"duration,omitempty"`
	Acute    *int32             `bson:"acute,omitempty"`
	Subacute *int32             `bson:"subacute,omitempty"`
	Chronic  *int32             `bson:"chronic,omitempty"`
	Symptom  []string           `bson:"symptom,omitempty"`
	Advice   *string            `bson:"advice,omitempty"`
	Question string             `bson:"question,omitempty"`
}

type SymptomCreateInput struct {
	Code     string   `bson:"code,omitempty"`
	Name     string   `bson:"name,omitempty"`
	Location *string  `bson:"location,omitempty"`
	Duration *int32   `bson:"duration,omitempty"`
	Acute    *int32   `bson:"acute,omitempty"`
	Subacute *int32   `bson:"subacute,omitempty"`
	Chronic  *int32   `bson:"chronic,omitempty"`
	Symptom  []string `bson:"symptom,omitempty"`
	Advice   *string  `bson:"advice,omitempty"`
	Question string   `bson:"question,omitempty"`
}

type SymptomUpdateInput struct {
	ID       string    `bson:"_id"`
	Code     *string   `bson:"code,omitempty"`
	Name     *string   `bson:"name,omitempty"`
	Location *string   `bson:"location,omitempty"`
	Duration *int32    `bson:"duration,omitempty"`
	Acute    *int32    `bson:"acute,omitempty"`
	Subacute *int32    `bson:"subacute,omitempty"`
	Chronic  *int32    `bson:"chronic,omitempty"`
	Symptom  *[]string `bson:"symptom,omitempty"`
	Advice   *string   `bson:"advice,omitempty"`
	Question *string   `bson:"question,omitempty"`
}
