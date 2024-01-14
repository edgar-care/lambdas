package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rdv struct {
	ID                primitive.ObjectID `bson:"_id"`
	DoctorID          string             `bson:"doctor_id,omitempty"`
	StartDate         int32              `bson:"start_date,omitempty"`
	EndDate           int32              `bson:"end_date,omitempty"`
	IdPatient         string             `bson:"id_patient,omitempty"`
	CancelationReason *string            `bson:"cancelation_reason,omitempty"`
}

type RdvCreateInput struct {
	DoctorID  string `bson:"doctor_id,omitempty"`
	StartDate int32  `bson:"start_date,omitempty"`
	EndDate   int32  `bson:"end_date,omitempty"`
	IdPatient string `bson:"id_patient,omitempty"`
}

type RdvUpdateInput struct {
	ID                string  `bson:"_id"`
	DoctorID          *string `bson:"doctor_id,omitempty"`
	StartDate         *int32  `bson:"start_date,omitempty"`
	EndDate           *int32  `bson:"end_date,omitempty"`
	IdPatient         *string `bson:"id_patient,omitempty"`
	CancelationReason *string `bson:"cancelation_reason,omitempty"`
}
