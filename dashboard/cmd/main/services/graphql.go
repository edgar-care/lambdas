package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Patient struct {
	Id                 string `json:"id"`
	OnboardingInfoID   string `json:"onboarding_info_id"`
	OnboardingHealthID string `json:"onboarding_health_id"`
}

type PatientOutput struct {
	Id                 *string `json:"id"`
	OnboardingInfoID   *string `json:"onboarding_info_id"`
	OnboardingHealthID *string `json:"onboarding_health_id"`
}

type Doctor struct {
	Id         string   `json:"id"`
	PatientIds []string `json:"patient_ids"`
}

type DoctorInput struct {
	Id         string   `json:"id"`
	PatientIds []string `json:"patient_ids"`
}

type DoctorOutput struct {
	Id         *string   `json:"id"`
	PatientIds *[]string `json:"patient_ids"`
}

type getPatientByIdResponse struct {
	Content PatientOutput `json:"getPatientById"`
}

type getAllPatientDoctorResponse struct {
	Content DoctorOutput `json:"getDoctorById"`
}

/*************** Implementations *****************/

func GetPatientById(id string) (Patient, error) {
	var patient getPatientByIdResponse
	var resp Patient
	query := `query getPatientById($id: String!) {
                getPatientById(id: $id) {
                    id,
					onboarding_info_id,
					onboarding_health_id
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetAllPatientDoctor(id string) ([]Doctor, error) {
	var allpatient getAllPatientDoctorResponse
	var resp []Doctor
	query := `query getDoctorById($id: String!){
                getDoctorById(id: $id) {
                    id,
					patient_ids
                }
            }`
	err := Query(query, map[string]interface{}{
		"id": id,
	}, &allpatient)
	_ = copier.Copy(&resp, &allpatient.Content)
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
