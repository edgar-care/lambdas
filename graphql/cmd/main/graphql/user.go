package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type userResolver struct {
	u *models.User
}

func (u *userResolver) ID() graphql.ID {
	return graphql.ID(u.u.ID.Hex())
}

func (u *userResolver) Email() string {
	return u.u.Email
}

func (u *userResolver) Password() string {
	return u.u.Password
}

func (u *userResolver) Name() *string {
	return u.u.Name
}

func (u *userResolver) Age() *int32 {
	return u.u.Age
}

func resolverFromUser(user *models.User) userResolver {
	return userResolver{u: user}
}

func (*Resolver) GetUsers() (*[]*userResolver, error) {
	users, err := db.FindUsers()
	lib.CheckError(err)

	var entities []*userResolver
	for i := range *users {
		resolv := resolverFromUser(&(*users)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) FindUser(args struct{ Id string }) (*userResolver, error) {
	user, err := db.FindUserByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromUser(user)
	return &resolv, nil
}

func (*Resolver) CreateUser(input models.UserCreateInput) (*userResolver, error) {
	var user = models.UserCreateInput{
		Name:     input.Name,
		Password: input.Password, //TODO crypt that
		Email:    input.Email,
		Age:      input.Age,
	}
	entity, err := db.InsertUser(&user)
	lib.CheckError(err)
	return &userResolver{entity}, nil
}

func (*Resolver) UpdateUser(input models.UserUpdateInput) (*userResolver, error) {
	res, err := db.UpdateUser(&input)
	lib.CheckError(err)
	return &userResolver{res}, nil
}

func (*Resolver) DeleteUser(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteUser(args.Id)
	lib.CheckError(err)
	return &result, err
}
