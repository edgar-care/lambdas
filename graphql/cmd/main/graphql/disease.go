package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type diseaseResolver struct {
	p *models.Disease
}

func (u *diseaseResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *diseaseResolver) Code() string {
	return u.p.Code
}

func (u *diseaseResolver) Name() string {
	return u.p.Name
}

func (u *diseaseResolver) Symptoms() []string {
	return u.p.Symptoms
}

func (u *diseaseResolver) Advice() *string {
	return u.p.Advice
}

func resolverFromDisease(p *models.Disease) diseaseResolver {
	return diseaseResolver{p: p}
}

func (*Resolver) GetDiseases() (*[]*diseaseResolver, error) {
	diseases, err := db.GetDiseases()
	lib.CheckError(err)

	var entities []*diseaseResolver
	for i := range *diseases {
		resolv := resolverFromDisease(&(*diseases)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDiseaseById(args struct{ Id string }) (*diseaseResolver, error) {
	disease, err := db.GetDiseaseByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromDisease(disease)
	return &resolv, nil
}

func (*Resolver) CreateDisease(disease models.DiseaseCreateInput) (*diseaseResolver, error) {
	entity, err := db.InsertDisease(&disease)
	lib.CheckError(err)

	return &diseaseResolver{entity}, nil
}

func (*Resolver) UpdateDisease(input models.DiseaseUpdateInput) (*diseaseResolver, error) {
	res, err := db.UpdateDisease(&input)
	lib.CheckError(err)

	return &diseaseResolver{res}, nil
}

func (*Resolver) DeleteDisease(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteDisease(args.Id)
	lib.CheckError(err)

	return &result, err
}
