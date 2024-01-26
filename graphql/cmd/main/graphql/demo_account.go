package graphql

import (
	"errors"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type demoAccountResolver struct {
	p *models.DemoAccount
}

func (u *demoAccountResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *demoAccountResolver) Email() string {
	return u.p.Email
}

func (u *demoAccountResolver) Password() string {
	return u.p.Password
}

func resolverFromDemoAccount(p *models.DemoAccount) demoAccountResolver {
	return demoAccountResolver{p: p}
}

func (*Resolver) GetDemoAccounts() (*[]*demoAccountResolver, error) {
	demoAccounts, err := db.GetDemoAccounts()
	lib.CheckError(err)

	var entities []*demoAccountResolver
	for i := range *demoAccounts {
		resolv := resolverFromDemoAccount(&(*demoAccounts)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDemoAccountById(args struct{ Id string }) (*demoAccountResolver, error) {
	demoAccount, err := db.GetDemoAccountByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromDemoAccount(demoAccount)
	return &resolv, nil
}

func (*Resolver) GetDemoAccountByEmail(args struct{ Email string }) (*demoAccountResolver, error) {
	demoAccount, err := db.GetDemoAccountByEmail(args.Email)
	lib.CheckError(err)

	resolv := resolverFromDemoAccount(demoAccount)
	return &resolv, nil
}

func (*Resolver) CreateDemoAccount(demoAccount models.DemoAccountCreateInput) (*demoAccountResolver, error) {
	_, err := db.GetDemoAccountByEmail(demoAccount.Email)
	if err == nil {
		return nil, errors.New("Email already exists")
	}
	entity, err := db.InsertDemoAccount(&demoAccount)
	lib.CheckError(err)

	return &demoAccountResolver{entity}, nil
}

func (*Resolver) UpdateDemoAccount(input models.DemoAccountUpdateInput) (*demoAccountResolver, error) {
	res, err := db.UpdateDemoAccount(&input)
	lib.CheckError(err)

	return &demoAccountResolver{res}, nil
}

func (*Resolver) DeleteDemoAccount(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteDemoAccount(args.Id)
	lib.CheckError(err)

	return &result, err
}
