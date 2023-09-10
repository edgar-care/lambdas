package graphql

import (
	"errors"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type adminResolver struct {
	p *models.Admin
}

func (u *adminResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *adminResolver) Email() string {
	return u.p.Email
}

func (u *adminResolver) Password() string {
	return u.p.Password
}

func (u *adminResolver) Name() string {
	return u.p.Name
}

func (u *adminResolver) LastName() string {
	return u.p.LastName
}

func resolverFromAdmin(p *models.Admin) adminResolver {
	return adminResolver{p: p}
}

func (*Resolver) GetAdmins() (*[]*adminResolver, error) {
	admins, err := db.GetAdmins()
	lib.CheckError(err)

	var entities []*adminResolver
	for i := range *admins {
		resolv := resolverFromAdmin(&(*admins)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetAdminById(args struct{ Id string }) (*adminResolver, error) {
	admin, err := db.GetAdminByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromAdmin(admin)
	return &resolv, nil
}

func (*Resolver) GetAdminByEmail(args struct{ Email string }) (*adminResolver, error) {
	admin, err := db.GetAdminByEmail(args.Email)
	lib.CheckError(err)

	resolv := resolverFromAdmin(admin)
	return &resolv, nil
}

func (*Resolver) CreateAdmin(admin models.AdminCreateInput) (*adminResolver, error) {
	_, err := db.GetAdminByEmail(admin.Email)
	if err == nil {
		return nil, errors.New("Email already exists")
	}
	entity, err := db.InsertAdmin(&admin)
	lib.CheckError(err)

	return &adminResolver{entity}, nil
}

func (*Resolver) UpdateAdmin(input models.AdminUpdateInput) (*adminResolver, error) {
	res, err := db.UpdateAdmin(&input)
	lib.CheckError(err)

	return &adminResolver{res}, nil
}

func (*Resolver) DeleteAdmin(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteAdmin(args.Id)
	lib.CheckError(err)

	return &result, err
}
