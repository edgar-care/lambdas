package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
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

type Symptom struct {
	ID       string   `json:"id"`
	Code     string   `json:"code"`
	Symptom  []string `json:"symptom"`
	Advice   string   `json:"advice"`
	Question string   `json:"question"`
}

type AnteChir struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Localisation    string    `json:"localisation"`
	InducedSymptoms []Symptom `json:"inducedsymptoms"`
}

type AnteChirOutput struct {
	ID              *string   `json:"id"`
	Name            *string   `json:"name"`
	Localisation    *string   `json:"localisation"`
	InducedSymptoms []Symptom `json:"inducedsymptoms"`
}

type AnteChirInput struct {
	Name            string    `json:"name"`
	Localisation    string    `json:"localisation"`
	InducedSymptoms []Symptom `json:"inducedsymptoms"`
}

type AnteDisease struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Chronicity float32     `json:"chronicity"`
	Chir       AnteChir    `json:"chir"`
	Treatment  []Treatment `json:"treatment"`
	Symptoms   []Symptom   `json:"symptoms"`
}

type AnteDiseaseOutput struct {
	ID         *string      `json:"id"`
	Name       *string      `json:"name"`
	Chronicity *float32     `json:"chronicity"`
	Chir       *AnteChir    `json:"chir"`
	Treatment  *[]Treatment `json:"treatment"`
	Symptoms   *[]Symptom   `json:"symptoms"`
}

type AnteDiseaseInput struct {
	Name       string      `json:"name"`
	Chronicity float32     `json:"chronicity"`
	Chir       AnteChir    `json:"chir"`
	Treatment  []Treatment `json:"treatment"`
	Symptoms   []Symptom   `json:"symptoms"`
}

type AnteFamily struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Disease []Disease `json:"disease"`
}

type AnteFamilyOutput struct {
	ID      *string   `json:"id"`
	Name    *string   `json:"name"`
	Disease []Disease `json:"disease"`
}

type AnteFamilyInput struct {
	Name    string    `json:"name"`
	Disease []Disease `json:"disease"`
}

type Treatment struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Disease     Disease   `json:"disease"`
	Symptoms    []Symptom `json:"symptoms"`
	SideEffects []Symptom `json:"sideeffects"`
}

type TreatmentOutput struct {
	ID          *string   `json:"id"`
	Name        *string   `json:"name"`
	Disease     *Disease  `json:"disease"`
	Symptoms    []Symptom `json:"symptoms"`
	SideEffects []Symptom `json:"sideeffects"`
}

type TreatmentInput struct {
	Name        string    `json:"name"`
	Disease     Disease   `json:"disease"`
	Symptoms    []Symptom `json:"symptoms"`
	SideEffects []Symptom `json:"sideeffects"`
}

/**************** GraphQL types *****************/

type getAnteChirByIdResponse struct {
	Content AnteChirOutput `json:"getAnteChirById"`
}

type createAnteChirResponse struct {
	Content AnteChirOutput `json:"createAnteChir"`
}

type getAnteDiseaseByIdResponse struct {
	Content AnteDiseaseOutput `json:"getAnteDiseaseById"`
}

type createAnteDiseaseResponse struct {
	Content AnteDiseaseOutput `json:"createAnteDisease"`
}

type getAnteFamilyByIdResponse struct {
	Content AnteFamilyOutput `json:"getAnteFamilyById"`
}

type createAnteFamilyResponse struct {
	Content AnteFamilyOutput `json:"createAnteFamily"`
}

type getTreatmentIdResponse struct {
	Content TreatmentOutput `json:"getTreatmentById"`
}

type createTreatmentResponse struct {
	Content TreatmentOutput `json:"createTreatment"`
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

func GetAnteChirById(id string) (AnteChir, error) {
	var antechir getAnteChirByIdResponse
	var resp AnteChir
	query := `query getAnteChirById($id: String!) {
                getAnteChirById(id: $id) {
                    id,
                    name,
                    localisation,
					inducedsymptoms,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &antechir)
	_ = copier.Copy(&resp, &antechir.Content)
	return resp, err
}

func CreateAnteChir(newAnteChir AnteChirInput) (AnteChir, error) {
	var antechir createAnteChirResponse
	var resp AnteChir
	query := `mutation createAnteChir($name: String!, $localisation: String!, $inducedsymptoms: [SymptomInput!]) {
        createAnteChir(name:$name, localisation:$localisation, inducedsymptoms:$inducedsymptoms) {
                    id,
					name,
                    localisation,
                    inducedsymptoms,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":            newAnteChir.Name,
		"localisation":    newAnteChir.Localisation,
		"inducedsymptoms": newAnteChir.InducedSymptoms,
	}, &antechir)
	_ = copier.Copy(&resp, &antechir.Content)
	return resp, err
}

func GetAnteDiseaseById(id string) (AnteDisease, error) {
	var antedisease getAnteDiseaseByIdResponse
	var resp AnteDisease
	query := `query getAnteDiseaseById($id: String!) {
                getAnteDiseaseById(id: $id) {
                    id,
                    name,
                    chronicity,
					chir,
					treatment,
					symptoms,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &antedisease)
	_ = copier.Copy(&resp, &antedisease.Content)
	return resp, err
}

func CreateAnteDisease(newAnteDisease AnteDiseaseInput) (AnteDisease, error) {
	var antedisease createAnteDiseaseResponse
	var resp AnteDisease
	query := `mutation createAnteDisease($name: String!, $chronicity: String!, $chir: AnteChirInput, $treatment: [TreatmentInput!], $symptoms: [SymptomInput!]!) {
        createAnteDisease(name:$name, localisation:$localisation, inducedsymptoms:$inducedsymptoms) {
                    id,
                    name,
                    chronicity,
					chir,
					treatment,
					symptoms,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":       newAnteDisease.Name,
		"chronicity": newAnteDisease.Chronicity,
		"chir":       newAnteDisease.Chir,
		"treatment":  newAnteDisease.Treatment,
		"symptoms":   newAnteDisease.Symptoms,
	}, &antedisease)
	_ = copier.Copy(&resp, &antedisease.Content)
	return resp, err
}

func GetAnteFamilyById(id string) (AnteFamily, error) {
	var antefamily getAnteFamilyByIdResponse
	var resp AnteFamily
	query := `query getAnteFamilyById($id: String!) {
                getAnteFamilyById(id: $id) {
                    id,
                    name,
                    disease,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &antefamily)
	_ = copier.Copy(&resp, &antefamily.Content)
	return resp, err
}

func CreateAnteFamily(newAnteFamily AnteFamilyInput) (AnteFamily, error) {
	var antefamily createAnteFamilyResponse
	var resp AnteFamily
	query := `mutation createAnteFamily($name: String!, $disease: [DiseaseInput!]!) {
        createAnteFamily(name:$name, disease:$disease) {
                    id,
                    name,
                    disease,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":    newAnteFamily.Name,
		"disease": newAnteFamily.Disease,
	}, &antefamily)
	_ = copier.Copy(&resp, &antefamily.Content)
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
