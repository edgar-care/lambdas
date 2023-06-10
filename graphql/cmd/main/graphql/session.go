package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type sessionResolver struct {
	p *models.Session
}

func (u *sessionResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *sessionResolver) Symptoms() []string {
	return u.p.Symptoms
}

func (u *sessionResolver) Age() int32 {
	return u.p.Age
}

func (u *sessionResolver) Height() int32 {
	return u.p.Height
}

func (u *sessionResolver) Weight() int32 {
	return u.p.Weight
}

func (u *sessionResolver) Sex() string {
	return u.p.Sex
}

func (u *sessionResolver) LastQuestion() string {
	return u.p.LastQuestion
}

func resolverFromSession(p *models.Session) sessionResolver {
	return sessionResolver{p: p}
}

func (*Resolver) GetSessions() (*[]*sessionResolver, error) {
	sessions, err := db.GetSessions()
	lib.CheckError(err)

	var entities []*sessionResolver
	for i := range *sessions {
		resolv := resolverFromSession(&(*sessions)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetSessionById(args struct{ Id string }) (*sessionResolver, error) {
	session, err := db.GetSessionByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromSession(session)
	return &resolv, nil
}

func (*Resolver) CreateSession(session models.SessionCreateInput) (*sessionResolver, error) {
	entity, err := db.InsertSession(&session)
	lib.CheckError(err)

	return &sessionResolver{entity}, nil
}

func (*Resolver) UpdateSession(input models.SessionUpdateInput) (*sessionResolver, error) {
	res, err := db.UpdateSession(&input)
	lib.CheckError(err)

	return &sessionResolver{res}, nil
}

func (*Resolver) DeleteSession(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteSession(args.Id)
	lib.CheckError(err)

	return &result, err
}
