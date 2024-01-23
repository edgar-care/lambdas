package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type alertResolver struct {
	p *models.Alert
}

func (u *alertResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *alertResolver) Name() string {
	return u.p.Name
}

func (u *alertResolver) Sex() *string {
	return u.p.Sex
}

func (u *alertResolver) Height() *int32 {
	return u.p.Height
}

func (u *alertResolver) Weight() *int32 {
	return u.p.Weight
}

func (u *alertResolver) Symptoms() []string {
	return u.p.Symptoms
}

func (u *alertResolver) Comment() string {
	return u.p.Comment
}

func resolverFromAlert(p *models.Alert) alertResolver {
	return alertResolver{p: p}
}

func (*Resolver) GetAlerts() (*[]*alertResolver, error) {
	alerts, err := db.GetAlerts()
	lib.CheckError(err)

	var entities []*alertResolver
	for i := range *alerts {
		resolv := resolverFromAlert(&(*alerts)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetAlertById(args struct{ Id string }) (*alertResolver, error) {
	alert, err := db.GetAlertByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromAlert(alert)
	return &resolv, nil
}

func (*Resolver) CreateAlert(alert models.AlertCreateInput) (*alertResolver, error) {
	entity, err := db.InsertAlert(&alert)
	lib.CheckError(err)

	return &alertResolver{entity}, nil
}

func (*Resolver) UpdateAlert(input models.AlertUpdateInput) (*alertResolver, error) {
	res, err := db.UpdateAlert(&input)
	lib.CheckError(err)

	return &alertResolver{res}, nil
}

func (*Resolver) DeleteAlert(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteAlert(args.Id)
	lib.CheckError(err)

	return &result, err
}
