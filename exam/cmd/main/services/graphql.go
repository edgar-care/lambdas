package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"os"

	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Disease struct {
	Id       string   `json:"id"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Symptoms []string `json:"symptoms"`
	Advice   *string  `json:"advice"`
}

type DiseaseOutput struct {
	Id       string    `json:"id"`
	Code     *string   `json:"code"`
	Name     *string   `json:"name"`
	Symptoms *[]string `json:"symptoms"`
	Advice   *string   `json:"advice"`
}

/**************** GraphQL types *****************/

type getDiseasesResponse struct {
	Content []DiseaseOutput `json:"getDiseases"`
}

/*************** Implementations *****************/

func GetDiseases() ([]Disease, error) {
	var diseases getDiseasesResponse
	var resp []Disease
	query := `query getDiseases {
                getDiseases {
                    id
                    code
                    name
                    symptoms
                    advice
                }
            }`
	err := Query(query, map[string]interface{}{}, &diseases)
	if err != nil {
		panic(err)
	}
	fmt.Print(diseases)
	_ = copier.Copy(&resp, &diseases.Content)
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
