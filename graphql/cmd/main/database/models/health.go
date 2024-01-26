package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Health struct {
	ID           primitive.ObjectID `bson:"_id"`
	PatientsAllergies    *[]string    `bson:"patients_allergies,omitempty"`
	PatientsIllness		*[]string    `bson:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`bson:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string	`bson:"patients_primary_doctor,omitempty"`
}

type HealthCreateInput struct {
	PatientsAllergies	*[]string	`bson:"patients_allergies,omitempty"`
	PatientsIllness		*[]string   	`bson:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`bson:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string	`bson:"patients_primary_doctor,omitempty"`
}

type HealthUpdateInput struct {
	ID           string    `bson:"_id"`
	PatientsAllergies	*[]string	`bson:"patients_allergies,omitempty"`
	PatientsIllness 	*[]string   `bson:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`bson:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor *string	`bson:"patients_primary_doctor,omitempty"`
}