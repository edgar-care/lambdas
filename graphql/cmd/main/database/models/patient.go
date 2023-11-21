package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID               primitive.ObjectID `bson:"_id"`
	Email            string             `bson:"email,omitempty"`
	Password         string             `bson:"password,omitempty"`
	OnboardingInfoID *string            `bson:"onboarding_info_id"`
	RendezVousIDs    *[]*string         `bson:"rendez_vous_ids"`
}

type PatientCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type PatientUpdateInput struct {
	ID               string     `bson:"_id"`
	Email            *string    `bson:"email,omitempty"`
	Password         *string    `bson:"password,omitempty"`
	OnboardingInfoID *string    `bson:"onboarding_info_id"`
	RendezVousIDs    *[]*string `bson:"rendez_vous_ids"`
}
