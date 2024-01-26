package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Doctor struct {
	ID            primitive.ObjectID `bson:"_id"`
	Email         string             `bson:"email,omitempty"`
	Password      string             `bson:"password,omitempty"`
	RendezVousIDs *[]*string         `bson:"rendez_vous_ids"`
	PatientIds    *[]*string         `bson:"patient_ids"`
	//SlotIDs       *[]*string         `bson:"slot_ids"`
}

type DoctorCreateInput struct {
	Email    string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

type DoctorUpdateInput struct {
	ID            string     `bson:"_id"`
	Email         *string    `bson:"email,omitempty"`
	Password      *string    `bson:"password,omitempty"`
	RendezVousIDs *[]*string `bson:"rendez_vous_ids"`
	PatientIds    *[]*string `bson:"patient_ids"`
	//SlotIDs       *[]*string `bson:"slot_ids"`
}
