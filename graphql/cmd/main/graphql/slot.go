package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type slotResolver struct {
	p *models.Slot
}

func (u *slotResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *slotResolver) DoctorID() string {
	return u.p.DoctorID
}

func (u *slotResolver) StartDate() int32 {
	return u.p.StartDate
}

func (u *slotResolver) AppointmentID() string {
	return u.p.AppointmentID
}

func (u *slotResolver) EndDate() int32 {
	return u.p.EndDate
}

func resolverFromSlot(p *models.Slot) slotResolver {
	return slotResolver{p: p}
}

func (*Resolver) GetDoctorSlot(args struct{ Doctor_id string }) (*[]*slotResolver, error) {
	slots, err := db.GetDoctorSlot(args.Doctor_id)
	lib.CheckError(err)

	var entities []*slotResolver
	for i := range *slots {
		resolv := resolverFromSlot(&(*slots)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetSlotById(args struct{ Id string }) (*slotResolver, error) {
	slot, err := db.GetSlotByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromSlot(slot)
	return &resolv, nil
}

func (*Resolver) CreateSlot(slot models.SlotCreateInput) (*slotResolver, error) {
	entity, err := db.InsertSlot(&slot)
	lib.CheckError(err)

	return &slotResolver{entity}, nil
}

func (*Resolver) UpdateSlot(input models.SlotUpdateInput) (*slotResolver, error) {
	res, err := db.UpdateSlot(&input)
	lib.CheckError(err)

	return &slotResolver{res}, nil
}

func (*Resolver) DeleteSlot(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteSlot(args.Id)
	lib.CheckError(err)

	return &result, err
}
