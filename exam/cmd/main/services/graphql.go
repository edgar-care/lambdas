package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/machinebox/graphql"
)

/********** Types ***********/

type DiseasesResponse struct {
	GetDiseases []Disease `json:"getDiseases"`
}

type Disease struct {
	Id       string   `json:"id"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	Symptoms []string `json:"symptoms"`
	Advice   string   `json:"advice"`
}

/*************** Implementations *****************/

func GetDiseases() ([]Disease, error) {
	var resp DiseasesResponse
	var err error

	jsonData := `{
		"getDiseases": [
			{
				"id": "65390b614c208161d18037ce",
				"code": "asthme",
				"name": "Asthme",
				"symptoms": [
					"respiration_difficile",
					"toux",
					"respiration_sifflante",
					"somnolence",
					"anxiete"
				],
				"advice": null
			},
			{
				"id": "65390c784c208161d18037cf",
				"code": "brulure_estomac",
				"name": "Brulure d'estomac",
				"symptoms": [
					"brulure_poitrine",
					"respiration_difficile",
					"boule_gorge",
					"toux"
				],
				"advice": null
			},
			{
				"id": "65392a0c467ff3023b4631f1",
				"code": "migraine_aura",
				"name": "Migraine avec aura",
				"symptoms": [
					"maux_de_tetes",
					"vision_trouble",
					"tache_visuel"
				],
				"advice": null
			},
			{
				"id": "65392a500ef9f8b48f6e0a7a",
				"code": "migraine_sans_aura",
				"name": "Migraine sans aura",
				"symptoms": [
					"maux_de_tetes",
					"douleur_pulsatile",
					"vomissements",
					"sensibilite_lumiere"
				],
				"advice": null
			},
			{
				"id": "65392a780ef9f8b48f6e0a7b",
				"code": "migraine_vestibulaire",
				"name": "Migraine vestibulaire",
				"symptoms": [
					"maux_de_tetes",
					"vertige",
					"perte_equilibre"
				],
				"advice": null
			}
		]
	}`

	if err = json.Unmarshal([]byte(jsonData), &resp); err != nil {
		fmt.Println("Error:", err)
	}
	return resp.GetDiseases, err
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
