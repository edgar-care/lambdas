package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Info struct {
	Id          string   `json:"id"`
	Name    	string `json:"name"`
	Surname 	string `json:"surname"`
	BirthDate 	string `json:"birthdate"`
	Sex 		string `json:"sex"`
	Weight 		int `json:"weight"`
	Height 		int `json:"height"`
}

type InfoOutput struct {
	Id           string    `json:"id"`
	Name    	*string `json:"name"`
	Surname 	*string `json:"surname"`
	BirthDate	*string `json:"birthdate"`
	Sex		*string `json:"sex"`
	Weight 		*int `json:"weight"`
	Height 		*int `json:"height"`
}

type InfoInput struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	BirthDate string `json:"birthdate"`
	Sex string `json:"sex"`
	Weight int `json:"weight"`
	Height int `json:"height"`
}

type Health struct {
	Id          string   `json:"id"`
	PatientsAllergies    *[]string `json:"patients_allergies,omitempty"`
	PatientsIllness 	*[]string `json:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string	`json:"patients_primary_doctor,omitempty"`
}

type HealthInput struct {
	PatientsAllergies	*[]string `json:"patients_allergies,omitempty"`
	PatientsIllness		*[]string `json:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string	`json:"patients_primary_doctor,omitempty"`
}

type HealthOutput struct {
	Id           string    `json:"id"`
	PatientsAllergies    *[]string `json:"patients_allergies,omitempty"`
	PatientsIllness 	*[]string `json:"patients_illness,omitempty"`
	PatientsTreatments	*[]string	`json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor *string	`json:"patients_primary_doctor,omitempty"`
}

type MedicalInfo struct {
	Info InfoInput `json:"onboarding_info"`
	Health HealthInput `json:"onboarding_health"`
}


type Patient struct {
	Id	string `json:"id"`
	OnboardingInfoID string `json:"onboarding_info_id"`
	OnboardingHealthID	string	`json:"onboarding_health_id"`
}

type PatientInput struct {
	Id	string `json:"id"`
	OnboardingInfoID string `json:"onboarding_info_id"`
	OnboardingHealthID	string	`json:"onboarding_health_id"`
}

type PatientOutput struct {
	Id       *string `json:"id"`
	OnboardingInfoID *string `json:"onboarding_info_id"`
	OnboardingHealthID	*string	`json:"onboarding_health_id"`
}


/**************** GraphQL types *****************/

type createInfoResponse struct {
	Content InfoOutput `json:"createInfo"`
}

type getInfoByIdResponse struct {
	Content InfoOutput `json:"getInfoById"`
}

type createHealthResponse struct {
	Content HealthOutput `json:"createHealth"`
}

type getHealthByIdResponse struct {
	Content HealthOutput `json:"getHealthById"`
}

type updateInfoResponse struct {
	Content InfoOutput `json:"updateInfo"`
}

type updateHealthResponse struct {
	Content HealthOutput `json:"updateHealth"`
}

type updateOnboardingResponse struct {
	Content PatientOutput `json:"updatePatient"`
}

type getPatientByIdResponse struct {
	Content PatientOutput `json:"getPatientById"`
}

/*************** Implementations *****************/

func CreateInfo(newInfo InfoInput) (Info, error) {
	var info createInfoResponse
	var resp Info
	query := `mutation createInfo($name: String!, $surname: String!, $birthdate: String!, $height: Int!, $weight: Int!, $sex: String!) {
            createInfo(name:$name, surname:$surname, birthdate:$birthdate, height:$height, weight:$weight, sex:$sex) {
                    id,
					name,
					birthdate,
					height,
					weight,
					sex,
					surname
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":      newInfo.Name,
		"birthdate":	newInfo.BirthDate,
		"height":        newInfo.Height,
		"weight":        newInfo.Weight,
		"sex":           newInfo.Sex,
		"surname": newInfo.Surname,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func GetInfoById(id string) (Info, error) {
	var info getInfoByIdResponse
	var resp Info
	query := `query getInfoById($id: String!) {
				getInfoById(id: $id) {
                    id,
					name,
					birthdate,
					height,
					weight,
					sex,
					surname
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func CreateHealth(newHealth HealthInput) (Health, error) {
	var health createHealthResponse
	var resp Health
	query := `mutation createHealth($patients_allergies: [String!], $patients_illness: [String!], $patients_treatments: [String!], $patients_primary_doctor: String!) {
            createHealth(patients_allergies:$patients_allergies, patients_illness:$patients_illness, patients_treatments:$patients_treatments, patients_primary_doctor:$patients_primary_doctor) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`
	err := Query(query, map[string]interface{}{
		"patients_allergies":		newHealth.PatientsAllergies,
		"patients_illness":		newHealth.PatientsIllness,
		"patients_treatments":	newHealth.PatientsTreatments,
		"patients_primary_doctor":	newHealth.PatientsPrimaryDoctor,
	}, &health)
	
	_ = copier.Copy(&resp, &health.Content)
	return resp, err
}

func GetHealthById(id string) (Health, error) {
	var health getHealthByIdResponse
	var resp Health
	query := `query getHealthById($id: String!) {
                getHealthById(id: $id) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`
	err := Query(query, map[string]interface{}{
		"id": id,
	}, &health)
	_ = copier.Copy(&resp, &health.Content)
	return resp, err
}

func UpdateMedicalInfo(id_info string, newInfo InfoInput) (Info, error) {
	var info updateInfoResponse
	var resp Info
	query := `mutation updateInfo($id: String! $name: String!, $surname: String!, $birthdate: String!, $height: Int!, $weight: Int!, $sex: String!) {
                updateInfo(id:$id, name:$name, surname:$surname, birthdate:$birthdate, height:$height, weight:$weight, sex:$sex) {
                    id,
					name,
					surname,
					birthdate,
					weight,
					sex,
					height
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id_info,
		"name": newInfo.Name,
		"surname": newInfo.Surname,
		"birthdate": newInfo.BirthDate,
		"weight": newInfo.Weight,
		"sex": newInfo.Sex,
		"height": newInfo.Height,

	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}


func UpdateMedicalHealth(id_health string, newHealth HealthInput) (Health, error) {
	var info updateHealthResponse
	var resp Health
	query := `mutation updateHealth($id: String!, $patients_allergies: [String!], $patients_illness: [String!], $patients_treatments: [String!], $patients_primary_doctor: String!) {
                updateHealth(id:$id, patients_allergies:$patients_allergies, patients_illness:$patients_illness, patients_treatments:$patients_treatments, patients_primary_doctor:$patients_primary_doctor) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id_health,
		"patients_allergies":		newHealth.PatientsAllergies,
		"patients_illness":		newHealth.PatientsIllness,
		"patients_treatments":	newHealth.PatientsTreatments,
		"patients_primary_doctor":	newHealth.PatientsPrimaryDoctor,

	}, &info)
	_ = copier.Copy(&resp, &info.Content)
	return resp, err
}

func AddOnboardingInfoID(updatePatient PatientInput) (Patient, error) {
	var patient updateOnboardingResponse
	var resp Patient
	query := `mutation updatePatient($id: String!, $onboarding_info_id: String) {
		updatePatient(id:$id, onboarding_info_id:$onboarding_info_id) {
                    id,
					onboarding_info_id
                }
            }`
	err := Query(query, map[string]interface{}{
		"id"	:	updatePatient.Id,
		"onboarding_info_id": updatePatient.OnboardingInfoID,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}


func AddOnboardingHealthID(updatePatient PatientInput) (Patient, error) {
	var patient updateOnboardingResponse
	var resp Patient
	query := `mutation updatePatient($id: String!, $onboarding_health_id: String) {
		updatePatient(id:$id, onboarding_health_id:$onboarding_health_id) {
                    id,
					onboarding_health_id
                }
            }`
	err := Query(query, map[string]interface{}{
		"id"	:	updatePatient.Id,
		"onboarding_health_id": updatePatient.OnboardingHealthID,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

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