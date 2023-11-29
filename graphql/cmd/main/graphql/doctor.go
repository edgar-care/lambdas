package graphql

import (
	"errors"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type doctorResolver struct {
	p *models.Doctor
}

func (u *doctorResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *doctorResolver) Email() string {
	return u.p.Email
}

func (u *doctorResolver) Password() string {
	return u.p.Password
}

func (u *doctorResolver) RendezVousIDs() *[]*string {
	return u.p.RendezVousIDs
}

func (r *doctorResolver) SlotIDs() *[]*string {
	return r.p.SlotIDs
}

func resolverFromDoctor(p *models.Doctor) doctorResolver {
	return doctorResolver{p: p}
}

func (*Resolver) GetDoctors() (*[]*doctorResolver, error) {
	doctors, err := db.GetDoctors()
	lib.CheckError(err)

	var entities []*doctorResolver
	for i := range *doctors {
		resolv := resolverFromDoctor(&(*doctors)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDoctorById(args struct{ Id string }) (*doctorResolver, error) {
	doctor, err := db.GetDoctorByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromDoctor(doctor)
	return &resolv, nil
}

func (*Resolver) GetDoctorByEmail(args struct{ Email string }) (*doctorResolver, error) {
	doctor, err := db.GetDoctorByEmail(args.Email)
	lib.CheckError(err)

	resolv := resolverFromDoctor(doctor)
	return &resolv, nil
}

func (*Resolver) CreateDoctor(doctor models.DoctorCreateInput) (*doctorResolver, error) {
	_, err := db.GetDoctorByEmail(doctor.Email)
	if err == nil {
		return nil, errors.New("Email already exists")
	}
	entity, err := db.InsertDoctor(&doctor)
	lib.CheckError(err)

	return &doctorResolver{entity}, nil
}

func (*Resolver) UpdateDoctor(input models.DoctorUpdateInput) (*doctorResolver, error) {
	res, err := db.UpdateDoctor(&input)
	lib.CheckError(err)

	return &doctorResolver{res}, nil
}

func (*Resolver) DeleteDoctor(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteDoctor(args.Id)
	lib.CheckError(err)

	return &result, err
}
