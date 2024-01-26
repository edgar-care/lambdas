package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type documentResolver struct {
	p *models.Document
}

func (u *documentResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *documentResolver) Name() string {
	return u.p.Name
}

func (u *documentResolver) OwnerID() string {
	return u.p.OwnerID
}

func (u *documentResolver) DocumentType() string {
	return u.p.DocumentType
}

func (u *documentResolver) Category() string {
	return u.p.Category
}

func (u *documentResolver) IsFavorite() bool {
	return u.p.IsFavorite
}

func (u *documentResolver) DownloadURL() string {
	return u.p.DownloadURL
}

func resolverFromDocument(p *models.Document) documentResolver {
	return documentResolver{p: p}
}

func (*Resolver) GetDocuments() (*[]*documentResolver, error) {
	documents, err := db.GetDocuments()
	lib.CheckError(err)

	var entities []*documentResolver
	for i := range *documents {
		resolv := resolverFromDocument(&(*documents)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetPatientDocument(args struct{ Id string }) (*[]*documentResolver, error) {
	documents, err := db.GetPatientDocument(args.Id)
	lib.CheckError(err)

	var entities []*documentResolver
	for i := range *documents {
		resolv := resolverFromDocument(&(*documents)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetDocumentById(args struct{ Id string }) (*documentResolver, error) {
	document, err := db.GetDocumentByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromDocument(document)
	return &resolv, nil
}

func (*Resolver) CreateDocument(document models.DocumentCreateInput) (*documentResolver, error) {
	entity, err := db.InsertDocument(&document)
	lib.CheckError(err)

	return &documentResolver{entity}, nil
}

func (*Resolver) UpdateDocument(input models.DocumentUpdateInput) (*documentResolver, error) {
	res, err := db.UpdateDocument(&input)
	lib.CheckError(err)

	return &documentResolver{res}, nil
}

func (*Resolver) DeleteDocument(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteDocument(args.Id)
	lib.CheckError(err)

	return &result, err
}
