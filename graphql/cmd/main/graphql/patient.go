package graphql

import (
	"errors"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type patientResolver struct {
	p *models.Patient
}

func (r *patientResolver) ID() graphql.ID {
	return graphql.ID(r.p.ID.Hex())
}

func (r *patientResolver) Email() string {
	return r.p.Email
}

func (r *patientResolver) Password() string {
	return r.p.Password
}

func (r *patientResolver) Name() string {
	return r.p.Name
}

func (r *patientResolver) Age() int32 {
	return r.p.Age
}

func (r *patientResolver) Height() int32 {
	return r.p.Height
}

func (r *patientResolver) Weight() int32 {
	return r.p.Weight
}

func (r *patientResolver) Sex() string {
	return r.p.Sex
}

func resolverFromPatient(p *models.Patient) patientResolver {
	return patientResolver{p: p}
}

func (*Resolver) GetPatients() (*[]*patientResolver, error) {
	patients, err := db.GetPatients()
	lib.CheckError(err)

	var entities []*patientResolver
	for i := range *patients {
		resolv := resolverFromPatient(&(*patients)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetPatientById(args struct{ Id string }) (*patientResolver, error) {
	patient, err := db.GetPatientByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromPatient(patient)
	return &resolv, nil
}

func (*Resolver) GetPatientByEmail(args struct{ Email string }) (*patientResolver, error) {
	patient, err := db.GetPatientByEmail(args.Email)
	lib.CheckError(err)

	resolv := resolverFromPatient(patient)
	return &resolv, nil
}

func (*Resolver) CreatePatient(patient models.PatientCreateInput) (*patientResolver, error) {
	if patient.Sex != "M" && patient.Sex != "F" {
		return nil, errors.New("Invalid value for Sex: " + patient.Sex)
	}
	_, err := db.GetPatientByEmail(patient.Email)
	if err == nil {
		return nil, errors.New("Email already exists")
	}
	entity, err := db.InsertPatient(&patient)
	lib.CheckError(err)

	return &patientResolver{entity}, nil
}

func (*Resolver) UpdatePatient(input models.PatientUpdateInput) (*patientResolver, error) {
	res, err := db.UpdatePatient(&input)
	lib.CheckError(err)

	return &patientResolver{res}, nil
}

func (*Resolver) DeletePatient(args struct{ Id string }) (*bool, error) {
	result, err := db.DeletePatient(args.Id)
	lib.CheckError(err)

	return &result, err
}
