package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

// SessionSymptom  ----------------------------------------------------------------------------------------------------
type sessionSymptomResolver struct {
	p models.SessionSymptom
}

func (u *sessionSymptomResolver) Name() string {
	return u.p.Name
}

func (u *sessionSymptomResolver) Presence() *bool {
	return u.p.Presence
}

func (u *sessionSymptomResolver) Duration() *int32 {
	return u.p.Duration
}

// Logs  --------------------------------------------------------------------------------------------------------------
type logsResolver struct {
	p models.Logs
}

func (u *logsResolver) Question() string {
	return u.p.Question
}

func (u *logsResolver) Answer() string {
	return u.p.Answer
}

// Session  -----------------------------------------------------------------------------------------------------------
type sessionResolver struct {
	p *models.Session
}

func (u *sessionResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *sessionResolver) Symptoms() []*sessionSymptomResolver {
	var SessionSymptomResolver []*sessionSymptomResolver
	if u.p.Symptoms == nil {
		return nil
	}
	for _, symptoms := range u.p.Symptoms {
		tmp := sessionSymptomResolver{p: symptoms}
		SessionSymptomResolver = append(SessionSymptomResolver, &tmp)
	}
	return SessionSymptomResolver
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

func (u *sessionResolver) AnteChirs() []string {
	return u.p.AnteChirs
}

func (u *sessionResolver) AnteDiseases() []string {
	return u.p.AnteDiseases
}

func (u *sessionResolver) Treatments() []string {
	return u.p.Treatments
}

func (u *sessionResolver) LastQuestion() string {
	return u.p.LastQuestion
}

func (u *sessionResolver) Logs() []*logsResolver {
	var LogsResolvers []*logsResolver
	if u.p.Logs == nil {
		return nil
	}
	for _, logs := range u.p.Logs {
		tmp := logsResolver{p: logs}
		LogsResolvers = append(LogsResolvers, &tmp)
	}
	return LogsResolvers
}

func (u *sessionResolver) Alerts() []string {
	return u.p.Alerts
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
