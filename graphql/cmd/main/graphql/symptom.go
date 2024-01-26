package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type symptomResolver struct {
	p *models.Symptom
}

func (u *symptomResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *symptomResolver) Code() string {
	return u.p.Code
}

func (u *symptomResolver) Name() string {
	return u.p.Name
}

func (u *symptomResolver) Location() *string {
	return u.p.Location
}

func (u *symptomResolver) Duration() *int32 {
	return u.p.Duration
}

func (u *symptomResolver) Acute() *int32 {
	return u.p.Acute
}

func (u *symptomResolver) Subacute() *int32 {
	return u.p.Subacute
}

func (u *symptomResolver) Chronic() *int32 {
	return u.p.Chronic
}

func (u *symptomResolver) Symptom() []string {
	return u.p.Symptom
}

func (u *symptomResolver) Advice() *string {
	return u.p.Advice
}

func (u *symptomResolver) Question() string {
	return u.p.Question
}

func resolverFromSymptom(p *models.Symptom) symptomResolver {
	return symptomResolver{p: p}
}

func (*Resolver) GetSymptoms() (*[]*symptomResolver, error) {
	symptoms, err := db.GetSymptoms()
	lib.CheckError(err)

	var entities []*symptomResolver
	for i := range *symptoms {
		resolv := resolverFromSymptom(&(*symptoms)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetSymptomById(args struct{ Id string }) (*symptomResolver, error) {
	symptom, err := db.GetSymptomByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromSymptom(symptom)
	return &resolv, nil
}

func (*Resolver) CreateSymptom(symptom models.SymptomCreateInput) (*symptomResolver, error) {
	entity, err := db.InsertSymptom(&symptom)
	lib.CheckError(err)

	return &symptomResolver{entity}, nil
}

func (*Resolver) UpdateSymptom(input models.SymptomUpdateInput) (*symptomResolver, error) {
	res, err := db.UpdateSymptom(&input)
	lib.CheckError(err)

	return &symptomResolver{res}, nil
}

func (*Resolver) DeleteSymptom(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteSymptom(args.Id)
	lib.CheckError(err)

	return &result, err
}
