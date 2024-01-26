package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type antefamilyResolver struct {
	p *models.AnteFamily
}

func (u *antefamilyResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *antefamilyResolver) Name() string {
	return u.p.Name
}

func (u *antefamilyResolver) Disease() []string {
	return u.p.Disease
}

func resolverFromAnteFamily(p *models.AnteFamily) antefamilyResolver {
	return antefamilyResolver{p: p}
}

func (*Resolver) GetAnteFamilies() (*[]*antefamilyResolver, error) {
	antefamilies, err := db.GetAnteFamilies()
	lib.CheckError(err)

	var entities []*antefamilyResolver
	for i := range *antefamilies {
		resolv := resolverFromAnteFamily(&(*antefamilies)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetAnteFamilyById(args struct{ Id string }) (*antefamilyResolver, error) {
	antefamily, err := db.GetAnteFamilyByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromAnteFamily(antefamily)
	return &resolv, nil
}

func (*Resolver) CreateAnteFamily(antefamily models.AnteFamilyCreateInput) (*antefamilyResolver, error) {
	entity, err := db.InsertAnteFamily(&antefamily)
	lib.CheckError(err)

	return &antefamilyResolver{entity}, nil
}

func (*Resolver) UpdateAnteFamily(input models.AnteFamilyUpdateInput) (*antefamilyResolver, error) {
	res, err := db.UpdateAnteFamily(&input)
	lib.CheckError(err)

	return &antefamilyResolver{res}, nil
}

func (*Resolver) DeleteAnteFamily(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteAnteFamily(args.Id)
	lib.CheckError(err)

	return &result, err
}
