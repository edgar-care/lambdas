package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type antediseaseResolver struct {
	p *models.AnteDisease
}

func (u *antediseaseResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *antediseaseResolver) Name() string {
	return u.p.Name
}

func (u *antediseaseResolver) Chronicity() float64 {
	return u.p.Chronicity
}

func (u *antediseaseResolver) Chir() *string {
	return u.p.Chir
}

func (u *antediseaseResolver) Treatment() *[]string {
	return u.p.Treatment
}

func (u *antediseaseResolver) Symptoms() *[]string {
	return u.p.Symptoms
}

func resolverFromAnteDisease(p *models.AnteDisease) antediseaseResolver {
	return antediseaseResolver{p: p}
}

func (*Resolver) GetAnteDiseases() (*[]*antediseaseResolver, error) {
	antediseases, err := db.GetAnteDiseases()
	lib.CheckError(err)

	var entities []*antediseaseResolver
	for i := range *antediseases {
		resolv := resolverFromAnteDisease(&(*antediseases)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetAnteDiseaseById(args struct{ Id string }) (*antediseaseResolver, error) {
	antedisease, err := db.GetAnteDiseaseByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromAnteDisease(antedisease)
	return &resolv, nil
}

func (*Resolver) CreateAnteDisease(antedisease models.AnteDiseaseCreateInput) (*antediseaseResolver, error) {
	entity, err := db.InsertAnteDisease(&antedisease)
	lib.CheckError(err)

	return &antediseaseResolver{entity}, nil
}

func (*Resolver) UpdateAnteDisease(input models.AnteDiseaseUpdateInput) (*antediseaseResolver, error) {
	res, err := db.UpdateAnteDisease(&input)
	lib.CheckError(err)

	return &antediseaseResolver{res}, nil
}

func (*Resolver) DeleteAnteDisease(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteAnteDisease(args.Id)
	lib.CheckError(err)

	return &result, err
}
