package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type rdvResolver struct {
	p *models.Rdv
}

func (u *rdvResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *rdvResolver) DoctorID() string {
	return u.p.DoctorID
}

func (u *rdvResolver) StartDate() int32 {
	return u.p.StartDate
}

func (u *rdvResolver) IdPatient() string {
	return u.p.IdPatient
}

func (u *rdvResolver) EndDate() int32 {
	return u.p.EndDate
}

func (u *rdvResolver) CancelationReason() *string {
	return u.p.CancelationReason
}

func resolverFromRdv(p *models.Rdv) rdvResolver {
	return rdvResolver{p: p}
}

func (*Resolver) GetPatientRdv(args struct{ Id_patient string }) (*[]*rdvResolver, error) {
	rdvs, err := db.GetPatientRdv(args.Id_patient)
	lib.CheckError(err)

	var entities []*rdvResolver
	for i := range *rdvs {
		resolv := resolverFromRdv(&(*rdvs)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDoctorRdv(args struct{ Doctor_Id string }) (*[]*rdvResolver, error) {
	rdvs, err := db.GetDoctorRdv(args.Doctor_Id)
	lib.CheckError(err)

	var entities []*rdvResolver
	for i := range *rdvs {
		resolv := resolverFromRdv(&(*rdvs)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetRdvById(args struct{ Id string }) (*rdvResolver, error) {
	rdv, err := db.GetRdvByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromRdv(rdv)
	return &resolv, nil
}

func (*Resolver) CreateRdv(rdv models.RdvCreateInput) (*rdvResolver, error) {
	entity, err := db.InsertRdv(&rdv)
	lib.CheckError(err)

	return &rdvResolver{entity}, nil
}

func (*Resolver) UpdateRdv(input models.RdvUpdateInput) (*rdvResolver, error) {
	res, err := db.UpdateRdv(&input)
	lib.CheckError(err)

	return &rdvResolver{res}, nil
}

func (*Resolver) DeleteRdv(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteRdv(args.Id)
	lib.CheckError(err)

	return &result, err
}

func (*Resolver) DeleteSlot(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteSlot(args.Id)
	lib.CheckError(err)

	return &result, err
}
