package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type healthResolver struct {
	p *models.Health
}

func (u *healthResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *healthResolver) Patientallergies() string {
	return u.p.Patientallergies
}


func (u *healthResolver) Patientsillness() string {
	return u.p.Patientsillness
}

func resolverFromHealth(p *models.Health) healthResolver {
	return healthResolver{p: p}
}

func (*Resolver) GetHealths() (*[]*healthResolver, error) {
	healths, err := db.GetHealths()
	lib.CheckError(err)

	var entities []*healthResolver
	for i := range *healths {
		resolv := resolverFromHealth(&(*healths)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetHealthById(args struct{ Id string }) (*healthResolver, error) {
	health, err := db.GetHealthByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromHealth(health)
	return &resolv, nil
}

func (*Resolver) CreateHealth(health models.HealthCreateInput) (*healthResolver, error) {
	entity, err := db.InsertHealth(&health)
	lib.CheckError(err)

	return &healthResolver{entity}, nil
}

func (*Resolver) UpdateHealth(input models.HealthUpdateInput) (*healthResolver, error) {
	res, err := db.UpdateHealth(&input)
	lib.CheckError(err)

	return &healthResolver{res}, nil
}

func (*Resolver) DeleteHealth(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteHealth(args.Id)
	lib.CheckError(err)

	return &result, err
}