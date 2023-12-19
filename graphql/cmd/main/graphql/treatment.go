package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type treatmentResolver struct {
	p *models.Treatment
}

func (u *treatmentResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *treatmentResolver) Name() string {
	return u.p.Name
}

func (u *treatmentResolver) Disease() models.Disease {
	return u.p.Disease
}

func (u *treatmentResolver) Symptoms() []models.Symptom {
	return u.p.Symptoms
}

func (u *treatmentResolver) SideEffects() []models.Symptom {
	return u.p.SideEffects
}

func resolverFromTreatment(p *models.Treatment) treatmentResolver {
	return treatmentResolver{p: p}
}

func (*Resolver) GetTreatments() (*[]*treatmentResolver, error) {
	treatments, err := db.GetTreatments()
	lib.CheckError(err)

	var entities []*treatmentResolver
	for i := range *treatments {
		resolv := resolverFromTreatment(&(*treatments)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetTreatmentById(args struct{ Id string }) (*treatmentResolver, error) {
	treatment, err := db.GetTreatmentByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromTreatment(treatment)
	return &resolv, nil
}

func (*Resolver) CreateTreatment(treatment models.TreatmentCreateInput) (*treatmentResolver, error) {
	entity, err := db.InsertTreatment(&treatment)
	lib.CheckError(err)

	return &treatmentResolver{entity}, nil
}

func (*Resolver) UpdateTreatment(input models.TreatmentUpdateInput) (*treatmentResolver, error) {
	res, err := db.UpdateTreatment(&input)
	lib.CheckError(err)

	return &treatmentResolver{res}, nil
}

func (*Resolver) DeleteTreatment(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteTreatment(args.Id)
	lib.CheckError(err)

	return &result, err
}
