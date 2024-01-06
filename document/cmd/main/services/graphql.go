package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type DocumentType string

const (
	DocumentTypeXray         DocumentType = "XRAY"
	DocumentTypePrescription DocumentType = "PRESCRIPTION"
	DocumentTypeOther        DocumentType = "OTHER"
	DocumentTypeCertificate  DocumentType = "CERTIFICATE"
)

type Category string

const (
	CategoryGENERAL Category = "GENERAL"
	CategoryFinance Category = "FINANCE"
)

type Document struct {
	Id           string `json:"id"`
	OwnerID      string `json:"owner_id"`
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Category     string `json:"category"`
	IsFavorite   bool   `json:"is_favorite"`
	DownloadURL  string `json:"download_url"`
}

type DocumentOutput struct {
	Id           string  `json:"id"`
	OwnerID      *string `json:"owner_id,omitempty"`
	Name         *string `json:"name,omitempty"`
	DocumentType *string `json:"document_type,omitempty"`
	Category     *string `json:"category,omitempty"`
	IsFavorite   bool    `json:"is_favorite,omitempty"`
	DownloadURL  *string `json:"download_url,omitempty"`
}

type DocumentInput struct {
	OwnerID      string `json:"owner_id"`
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Category     string `json:"category"`
	IsFavorite   bool   `json:"is_favorite"`
	DownloadURL  string `json:"download_url"`
}

type Patient struct {
	Id          string   `json:"id"`
	DocumentIDs []string `json:"document_ids"`
}

type PatientInput struct {
	Id          string   `json:"id"`
	DocumentIDs []string `json:"document_ids"`
}

type PatientOuput struct {
	Id         string    `json:"id"`
	DocumentID *[]string `json:"document_ids"`
}

/**************** GraphQL types *****************/

type uploadDocumentResponse struct {
	Content DocumentOutput `json:"createDocument"`
}

type getDocumentByIdResponse struct {
	Content DocumentOutput `json:"getDocumentById"`
}

type updateFavoriteByIdResponse struct {
	Content bool `json:"updateFavoriteById"`
}

type deleteDocumentResponse struct {
	Content DocumentOutput `json:"deleteDocument"`
}

type updatePatientResponse struct {
	Content PatientOuput `json:"updatePatient"`
}

/*************** Implementations *****************/

func CreateDocument(ownerID string, documentInput DocumentInput, downloadURL string) (Document, error) {
	var document uploadDocumentResponse
	var resp Document

	query := `
		mutation createDocument($owner_id: String!, $name: String!, $document_type: String!, $category: String!, $is_favorite: Boolean!, $download_url: String!) {
			createDocument(owner_id: $owner_id, name: $name, document_type: $document_type, category: $category, is_favorite: $is_favorite, download_url: $download_url) {
				id,
				owner_id,
				name,
				document_type,
				category,
				is_favorite,
				download_url
			}
		}
	`

	err := Query(query, map[string]interface{}{
		"owner_id":      ownerID,
		"name":          documentInput.Name,
		"document_type": documentInput.DocumentType,
		"category":      documentInput.Category,
		"is_favorite":   documentInput.IsFavorite,
		"download_url":  downloadURL,
	}, &document)
	_ = copier.Copy(&resp, &document.Content)

	return resp, err
}

func GetDocument(id string) (Document, error) {
	var document getDocumentByIdResponse
	var resp Document
	query := `query getDocumentById($id: String!) {
                getDocumentById(id: $id) {
                    id,
					owner_id,
					name,
					document_type,
					category,
					is_favorite
					download_url
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &document)
	_ = copier.Copy(&resp, &document.Content)
	return resp, err
}

func UpdateFavoriteById(id string, isFavorite bool) (Document, error) {
	var updateFavorite updateFavoriteByIdResponse
	var resp Document

	query := `
		mutation updateDocument($id: String!, $is_favorite: Boolean!) {
			updateDocument(id: $id, is_favorite: $is_favorite) {
				id,
				owner_id,
				name,
				document_type,
				category,
				is_favorite,
				download_url
			}
		}`
	err := Query(query, map[string]interface{}{
		"id":          id,
		"is_favorite": isFavorite,
	}, &updateFavorite)
	_ = copier.Copy(&resp, &updateFavorite.Content)

	return resp, err
}

func DeleteDocument(id string) (Document, error) {
	var deleteDocument deleteDocumentResponse
	var resp Document
	query := `mutation deleteDocument($id: String!) {
		deleteDocument(id:$id) {
                }
            }`
	err := Query(query, map[string]interface{}{
		"id": id,
	}, &deleteDocument)
	_ = copier.Copy(&resp, &deleteDocument.Content)
	return resp, err
}

/*************** Update Patient *****************/

func UpdatePatient(updatePatient PatientInput) (Patient, error) {
	var patient updatePatientResponse
	var resp Patient
	query := `mutation updatePatient($id: String!, $document_ids: [String]) {
		updatePatient(id:$id, document_ids:$document_ids) {
                    id,
					document_ids
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":           updatePatient.Id,
		"document_ids": updatePatient.DocumentIDs,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetPatientById(id string) (Patient, error) {
	var patient updatePatientResponse
	var resp Patient
	query := `query getPatientById($id: String!) {
                getPatientById(id: $id) {
                    id,
					document_ids
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func createClient() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL_URL"))
}

func Query(query string, variables map[string]interface{}, respData interface{}) error {
	var request = graphql.NewRequest(query)
	var ctx = context.Background()
	for key, value := range variables {
		request.Var(key, value)
	}
	request.Header.Set(os.Getenv("API_KEY"), os.Getenv("API_KEY_VALUE"))
	err := createClient().Run(ctx, request, respData)
	return err
}
