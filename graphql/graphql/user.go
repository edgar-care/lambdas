package graphql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type user struct {
	ID    graphql.ID
	email string
	name  *string
	age   *int32
}

type userCreateInput struct {
	Email string
	Name  *string
	Age   *int32
}

type userResolver struct {
	u *user
}

func (u *userResolver) ID() graphql.ID {
	return u.u.ID
}

func (u *userResolver) Email() string {
	return u.u.email
}

func (u *userResolver) Name() *string {
	return u.u.name
}

func (u *userResolver) Age() *int32 {
	return u.u.age
}

var names = []string{"John", "Paul", "Ringo"}
var ages = []int32{20, 30, 40}

var users = []*userResolver{
	{
		&user{
			ID:    graphql.ID("1"),
			email: "john@example.com",
			name:  &names[0],
			age:   &ages[0],
		},
	},
	{
		&user{
			ID:    graphql.ID("2"),
			email: "fabrice@gmail.com",
			age:   &ages[1],
		},
	},
	{
		&user{
			ID:    graphql.ID("3"),
			email: "test@test.com",
			name:  &names[2],
		},
	},
}

func (*Resolver) GetUsers() (*[]*userResolver, error) {
	return &users, nil
}

func (*Resolver) CreateUser(input userCreateInput) (*userResolver, error) {
	var user = user{
		ID:    graphql.ID(strconv.Itoa(len(users))),
		name:  input.Name,
		email: input.Email,
		age:   input.Age,
	}
	users = append(users, &userResolver{&user})
	return &userResolver{&user}, nil
}
