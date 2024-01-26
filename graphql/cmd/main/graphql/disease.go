package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type symptomWeightResolver struct {
	p *models.SymptomWeight
}

func (u *symptomWeightResolver) Key() string {
	return u.p.Key
}

func (u *symptomWeightResolver) Value() float64 {
	return u.p.Value
}

type diseaseResolver struct {
	p *models.Disease
}

func (u *diseaseResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *diseaseResolver) Code() string {
	return u.p.Code
}

func (u *diseaseResolver) Name() string {
	return u.p.Name
}

func (u *diseaseResolver) Symptoms() []string {
	return u.p.Symptoms
}

func (u *diseaseResolver) SymptomsAcute() *[]*symptomWeightResolver {
	var SymptomWeightResolvers []*symptomWeightResolver
	if u.p.SymptomsAcute == nil {
		return nil
	}
	for _, symptomWeight := range *u.p.SymptomsAcute {
		tmp := symptomWeightResolver{p: &symptomWeight}
		SymptomWeightResolvers = append(SymptomWeightResolvers, &tmp)
	}
	return &SymptomWeightResolvers
}

func (u *diseaseResolver) SymptomsSubacute() *[]*symptomWeightResolver {
	var SymptomWeightResolvers []*symptomWeightResolver
	if u.p.SymptomsSubacute == nil {
		return nil
	}
	for _, symptomWeight := range *u.p.SymptomsSubacute {
		tmp := symptomWeightResolver{p: &symptomWeight}
		SymptomWeightResolvers = append(SymptomWeightResolvers, &tmp)
	}
	return &SymptomWeightResolvers
}

func (u *diseaseResolver) SymptomsChronic() *[]*symptomWeightResolver {
	var SymptomWeightResolvers []*symptomWeightResolver
	if u.p.SymptomsChronic == nil {
		return nil
	}
	for _, symptomWeight := range *u.p.SymptomsChronic {
		tmp := symptomWeightResolver{p: &symptomWeight}
		SymptomWeightResolvers = append(SymptomWeightResolvers, &tmp)
	}
	return &SymptomWeightResolvers
}

func (u *diseaseResolver) Advice() *string {
	return u.p.Advice
}

func resolverFromDisease(p *models.Disease) diseaseResolver {
	return diseaseResolver{p: p}
}

func (*Resolver) GetDiseases() (*[]*diseaseResolver, error) {
	diseases, err := db.GetDiseases()
	lib.CheckError(err)

	var entities []*diseaseResolver
	for i := range *diseases {
		resolv := resolverFromDisease(&(*diseases)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDiseaseById(args struct{ Id string }) (*diseaseResolver, error) {
	disease, err := db.GetDiseaseByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromDisease(disease)
	return &resolv, nil
}

func (*Resolver) CreateDisease(disease models.DiseaseCreateInput) (*diseaseResolver, error) {
	entity, err := db.InsertDisease(&disease)
	lib.CheckError(err)

	return &diseaseResolver{entity}, nil
}

func (*Resolver) UpdateDisease(input models.DiseaseUpdateInput) (*diseaseResolver, error) {
	res, err := db.UpdateDisease(&input)
	lib.CheckError(err)

	return &diseaseResolver{res}, nil
}

func (*Resolver) DeleteDisease(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteDisease(args.Id)
	lib.CheckError(err)

	return &result, err
}
