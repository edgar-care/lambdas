package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type antechirResolver struct {
	p *models.AnteChir
}

func (u *antechirResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *antechirResolver) Name() string {
	return u.p.Name
}

func (u *antechirResolver) Localisation() string {
	return u.p.Localisation
}

func (u *antechirResolver) InducedSymptoms() *[]string {
	return u.p.InducedSymptoms
}

func resolverFromAnteChir(p *models.AnteChir) antechirResolver {
	return antechirResolver{p: p}
}

func (*Resolver) GetAnteChirs() (*[]*antechirResolver, error) {
	antechirs, err := db.GetAnteChirs()
	lib.CheckError(err)

	var entities []*antechirResolver
	for i := range *antechirs {
		resolv := resolverFromAnteChir(&(*antechirs)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetAnteChirById(args struct{ Id string }) (*antechirResolver, error) {
	antechir, err := db.GetAnteChirByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromAnteChir(antechir)
	return &resolv, nil
}

func (*Resolver) CreateAnteChir(antechir models.AnteChirCreateInput) (*antechirResolver, error) {
	entity, err := db.InsertAnteChir(&antechir)
	lib.CheckError(err)

	return &antechirResolver{entity}, nil
}

func (*Resolver) UpdateAnteChir(input models.AnteChirUpdateInput) (*antechirResolver, error) {
	res, err := db.UpdateAnteChir(&input)
	lib.CheckError(err)

	return &antechirResolver{res}, nil
}

func (*Resolver) DeleteAnteChir(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteAnteChir(args.Id)
	lib.CheckError(err)

	return &result, err
}
