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
	Age 		int `json:"age"`
	Sexe 		string `json:"sexe"`
	Weight 		string `json:"weight"`
	Height 		string `json:"height"`
}

type InfoOutput struct {
	Id           string    `json:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Age *int `json:"age"`
	Sexe *string `json:"sexe"`
	Weight *int `json:"weight"`
	Height *int `json:"height"`
}

type InfoInput struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
	Sexe string `json:"sexe"`
	Weight int `json:"weight"`
	Height int `json:"height"`
}

type Health struct {
	Id          string   `json:"id"`
	Patientallergies    string `json:"patientallergies,omitempty"`
	Patientsillness string `json:"patientsillness,omitempty"`
}

type HealthInput struct {
	Patientallergies    string `json:"patientallergies,omitempty"`
	Patientsillness string `json:"patientsillness,omitempty"`
}

type HealthOutput struct {
	Id           string    `json:"id"`
	Patientallergies    string `json:"patientallergies,omitempty"`
	Patientsillness string `json:"patientsillness,omitempty"`
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

/*************** Implementations *****************/

func CreateInfo(newInfo InfoInput) (Info, error) {
	var info createInfoResponse
	var resp Info
	query := `mutation createInfo($name: String!, $surname: String!, $age: Int!, $height: Int!, $weight: Int!, $sexe: String!) {
            createInfo(name:$name, surname:$surname, age:$age, height:$height, weight:$weight, sexe:$sexe) {
                    id,
					name,
					age,
					height,
					weight,
					sexe,
					surname
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":      newInfo.Name,
		"age":           newInfo.Age,
		"height":        newInfo.Height,
		"weight":        newInfo.Weight,
		"sexe":           newInfo.Sexe,
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
					age,
					height,
					weight,
					sexe,
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
	query := `mutation createHealth($patientallergies: String!, $patientsillness: String!) {
            createHealth(patientallergies:$patientallergies, patientsillness:$patientsillness) {
                    id,
					patientallergies,
					patientsillness
                }
            }`
	err := Query(query, map[string]interface{}{
		"patientallergies":      newHealth.Patientallergies,
		"patientsillness":           newHealth.Patientsillness,
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
					patientallergies,
					patientsillness
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &health)
	_ = copier.Copy(&resp, &health.Content)
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
	err := createClient().Run(ctx, request, respData)
	return err
}