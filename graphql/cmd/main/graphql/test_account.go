package graphql

import (
	"errors"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type testAccountResolver struct {
	p *models.TestAccount
}

func (u *testAccountResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *testAccountResolver) Email() string {
	return u.p.Email
}

func (u *testAccountResolver) Password() string {
	return u.p.Password
}

func resolverFromTestAccount(p *models.TestAccount) testAccountResolver {
	return testAccountResolver{p: p}
}

func (*Resolver) GetTestAccounts() (*[]*testAccountResolver, error) {
	testAccounts, err := db.GetTestAccounts()
	lib.CheckError(err)

	var entities []*testAccountResolver
	for i := range *testAccounts {
		resolv := resolverFromTestAccount(&(*testAccounts)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetTestAccountById(args struct{ Id string }) (*testAccountResolver, error) {
	testAccount, err := db.GetTestAccountByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromTestAccount(testAccount)
	return &resolv, nil
}

func (*Resolver) GetTestAccountByEmail(args struct{ Email string }) (*testAccountResolver, error) {
	testAccount, err := db.GetTestAccountByEmail(args.Email)
	lib.CheckError(err)

	resolv := resolverFromTestAccount(testAccount)
	return &resolv, nil
}

func (*Resolver) CreateTestAccount(testAccount models.TestAccountCreateInput) (*testAccountResolver, error) {
	_, err := db.GetTestAccountByEmail(testAccount.Email)
	if err == nil {
		return nil, errors.New("Email already exists")
	}
	entity, err := db.InsertTestAccount(&testAccount)
	lib.CheckError(err)

	return &testAccountResolver{entity}, nil
}

func (*Resolver) UpdateTestAccount(input models.TestAccountUpdateInput) (*testAccountResolver, error) {
	res, err := db.UpdateTestAccount(&input)
	lib.CheckError(err)

	return &testAccountResolver{res}, nil
}

func (*Resolver) DeleteTestAccount(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteTestAccount(args.Id)
	lib.CheckError(err)

	return &result, err
}
