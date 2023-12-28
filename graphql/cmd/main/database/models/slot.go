package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Slot struct {
	ID            primitive.ObjectID `bson:"_id"`
	DoctorID      string             `bson:"doctor_id,omitempty"`
	StartDate     int32              `bson:"start_date,omitempty"`
	EndDate       int32              `bson:"end_date,omitempty"`
	AppointmentID string             `bson:"appointment_id,omitempty"`
}

type SlotCreateInput struct {
	DoctorID      string `bson:"doctor_id,omitempty"`
	StartDate     int32  `bson:"start_date,omitempty"`
	EndDate       int32  `bson:"end_date,omitempty"`
	AppointmentID string `bson:"appointment_id,omitempty"`
}

type SlotUpdateInput struct {
	ID            string  `bson:"_id"`
	DoctorID      *string `bson:"doctor_id,omitempty"`
	StartDate     *int32  `bson:"start_date,omitempty"`
	EndDate       *int32  `bson:"end_date,omitempty"`
	AppointmentID *string `bson:"appointment_id,omitempty"`
}
