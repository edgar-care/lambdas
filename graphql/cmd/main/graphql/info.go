package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type infoResolver struct {
	p *models.Info
}

func (u *infoResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *infoResolver) Name() string {
	return u.p.Name
}

func (u *infoResolver) BirthDate() string {
	return u.p.BirthDate
}

func (u *infoResolver) Height() int32 {
	return u.p.Height
}

func (u *infoResolver) Weight() int32 {
	return u.p.Weight
}

func (u *infoResolver) Sex() string {
	return u.p.Sex
}

func (u *infoResolver) Surname() string {
	return u.p.Surname
}

func resolverFromInfo(p *models.Info) infoResolver {
	return infoResolver{p: p}
}

func (*Resolver) GetInfos() (*[]*infoResolver, error) {
	infos, err := db.GetInfos()
	lib.CheckError(err)

	var entities []*infoResolver
	for i := range *infos {
		resolv := resolverFromInfo(&(*infos)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetInfoById(args struct{ Id string }) (*infoResolver, error) {
	info, err := db.GetInfoByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromInfo(info)
	return &resolv, nil
}

func (*Resolver) CreateInfo(info models.InfoCreateInput) (*infoResolver, error) {
	entity, err := db.InsertInfo(&info)
	lib.CheckError(err)

	return &infoResolver{entity}, nil
}

func (*Resolver) UpdateInfo(input models.InfoUpdateInput) (*infoResolver, error) {
	res, err := db.UpdateInfo(&input)
	lib.CheckError(err)

	return &infoResolver{res}, nil
}

func (*Resolver) DeleteInfo(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteInfo(args.Id)
	lib.CheckError(err)

	return &result, err
}
