package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Rdv struct {
	Id        string `json:"id"`
	DoctorID  string `json:"doctor_id"`
	IdPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type RdvOutput struct {
	Id        string  `json:"id"`
	DoctorID  *string `json:"doctor_id"`
	IdPatient *string `json:"id_patient"`
	StartDate *int    `json:"start_date"`
	EndDate   *int    `json:"end_date"`
}

type RdvInput struct {
	Id        string `json:"id"`
	DoctorID  string `json:"doctor_id"`
	IdPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type Patient struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type PatientInput struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type PatientOutput struct {
	Id            *string   `json:"id"`
	RendezVousIDs *[]string `json:"rendez_vous_ids"`
}

type Doctor struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type DoctorInput struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type DoctorOutput struct {
	Id            *string   `json:"id"`
	RendezVousIDs *[]string `json:"rendez_vous_ids"`
}

/**************** GraphQL types *****************/

type updateRdvResponse struct {
	Content RdvOutput `json:"updateRdv"`
}

type getOneRdvByIdResponse struct {
	Content RdvOutput `json:"GetRdvById"`
}

type getAllRdvResponse struct {
	Content []RdvOutput `json:"getPatientRdv"`
}

type updatePatientResponse struct {
	Content PatientOutput `json:"updatePatient"`
}

type deleteRdvByIdResponse struct {
	Content RdvOutput `json:"DeleteRdvById"`
}

type getRdvDoctorResponse struct {
	Content []RdvOutput `json:"getDoctorRdv"`
}

/*************** Implementations *****************/

func UpdateRdv(id_patient string, rdv_id string) (Rdv, error) {
	var rdv updateRdvResponse
	var resp Rdv
	query := `mutation updateRdv($id: String!, $id_patient: String!) {
		updateRdv(id:$id, id_patient:$id_patient) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":         rdv_id,
		"id_patient": id_patient,
	}, &rdv)
	_ = copier.Copy(&resp, &rdv.Content)
	return resp, err
}

func GetRdvById(id string) (Rdv, error) {
	var onerdv getOneRdvByIdResponse
	var resp Rdv
	query := `query getRdvById($id: String!) {
                getRdvById(id: $id) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &onerdv)
	_ = copier.Copy(&resp, &onerdv.Content)
	return resp, err
}

func GetAllRdv(id string) ([]Rdv, error) {
	var allrdv getAllRdvResponse
	var resp []Rdv
	query := `query getPatientRdv($id_patient: String!){
                getPatientRdv(id_patient: $id_patient) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"id_patient": id,
	}, &allrdv)
	_ = copier.Copy(&resp, &allrdv.Content)
	return resp, err
}

// ======================================================================================== //

func GetRdvDoctorById(id string) ([]Rdv, error) {
	var allrdv getRdvDoctorResponse
	var resp []Rdv
	query := `query getDoctorRdv($doctor_id: String!){
                getDoctorRdv(doctor_id: $doctor_id) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"doctor_id": id,
	}, &allrdv)
	_ = copier.Copy(&resp, &allrdv.Content)
	return resp, err
}

// ============================================================================================== //
// Patient
func UpdatePatient(updatePatient PatientInput) (Patient, error) {
	var patient updatePatientResponse
	var resp Patient
	query := `mutation updatePatient($id: String!, $rendez_vous_ids: [String]) {
		updatePatient(id:$id, rendez_vous_ids:$rendez_vous_ids) {
                    id,
					rendez_vous_ids
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":              updatePatient.Id,
		"rendez_vous_ids": updatePatient.RendezVousIDs,
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
					rendez_vous_ids
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
