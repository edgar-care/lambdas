package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/copier"

	"github.com/machinebox/graphql"
)

/********** Types ***********/

type SymptomWeight struct {
	Key   string  `bson:"key"`
	Value float64 `bson:"value,omitempty"`
}

type Alert struct {
	ID       string   `json:"_id"`
	Name     string   `json:"name"`
	Sex      *string  `json:"sex"`
	Height   *int32   `json:"height"`
	Weight   *int32   `json:"weight"`
	Symptoms []string `json:"symptoms"`
	Comment  string   `json:"comment"`
}

type AlertOutput struct {
	ID       *string   `json:"_id"`
	Name     *string   `json:"name"`
	Sex      *string   `json:"sex"`
	Height   *int32    `json:"height"`
	Weight   *int32    `json:"weight"`
	Symptoms *[]string `json:"symptoms"`
	Comment  *string   `json:"comment"`
}

type AlertInput struct {
	ID       string   `json:"_id"`
	Name     string   `json:"name"`
	Sex      *string  `json:"sex"`
	Height   *int32   `json:"height"`
	Weight   *int32   `json:"weight"`
	Symptoms []string `json:"symptoms"`
	Comment  string   `json:"comment"`
}

type Disease struct {
	ID               string           `json:"id"`
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	Symptoms         []string         `json:"symptoms"`
	SymptomsAcute    *[]SymptomWeight `json:"symptoms_acute"`
	SymptomsSubacute *[]SymptomWeight `json:"symptoms_subacute"`
	SymptomsChronic  *[]SymptomWeight `json:"symptoms_chronic"`
	Advice           string           `json:"advice"`
}

type DiseaseOutput struct {
	ID               *string          `json:"id"`
	Code             *string          `json:"code"`
	Name             *string          `json:"name"`
	Symptoms         *[]string        `json:"symptoms"`
	SymptomsAcute    *[]SymptomWeight `json:"symptoms_acute"`
	SymptomsSubacute *[]SymptomWeight `json:"symptoms_subacute"`
	SymptomsChronic  *[]SymptomWeight `json:"symptoms_chronic"`
	Advice           *string          `json:"advice"`
}

type DiseaseInput struct {
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	Symptoms         []string         `json:"symptoms"`
	SymptomsAcute    *[]SymptomWeight `json:"symptoms_acute"`
	SymptomsSubacute *[]SymptomWeight `json:"symptoms_subacute"`
	SymptomsChronic  *[]SymptomWeight `json:"symptoms_chronic"`
	Advice           string           `json:"advice"`
}

type Symptom struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Duration int32  `json:"duration"`
	Acute    int32  `json:"acute"`
	Subacute int32  `json:"subacute"`
	Chronic  int32  `json:"chronic"`
	Advice   string `json:"advice"`
	Question string `json:"question"`
}

type SymptomOutput struct {
	ID       *string `json:"id"`
	Code     *string `json:"code"`
	Name     *string `json:"name"`
	Location *string `json:"location"`
	Duration int32   `json:"duration"`
	Acute    int32   `json:"acute"`
	Subacute int32   `json:"subacute"`
	Chronic  int32   `json:"chronic"`
	Advice   *string `json:"advice"`
	Question *string `json:"question"`
}

type SymptomInput struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Duration int32  `json:"duration"`
	Acute    int32  `json:"acute"`
	Subacute int32  `json:"subacute"`
	Chronic  int32  `json:"chronic"`
	Advice   string `json:"advice"`
	Question string `json:"question"`
}

type AnteChir struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Localisation    string   `json:"localisation"`
	InducedSymptoms []string `json:"induced_symptoms"`
}

type AnteChirOutput struct {
	ID              *string   `json:"id"`
	Name            *string   `json:"name"`
	Localisation    *string   `json:"localisation"`
	InducedSymptoms *[]string `json:"induced_symptoms"`
}

type AnteChirInput struct {
	Name            string   `json:"name"`
	Localisation    string   `json:"localisation"`
	InducedSymptoms []string `json:"induced_symptoms"`
}

type AnteDisease struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Chronicity float64  `json:"chronicity"`
	Chir       string   `json:"chir"`
	Treatment  []string `json:"treatment"`
	Symptoms   []string `json:"symptoms"`
}

type AnteDiseaseOutput struct {
	ID         *string   `json:"id"`
	Name       *string   `json:"name"`
	Chronicity *float64  `json:"chronicity"`
	Chir       *string   `json:"chir"`
	Treatment  *[]string `json:"treatment"`
	Symptoms   *[]string `json:"symptoms"`
}

type AnteDiseaseInput struct {
	Name       string   `json:"name"`
	Chronicity float64  `json:"chronicity"`
	Chir       string   `json:"chir"`
	Treatment  []string `json:"treatment"`
	Symptoms   []string `json:"symptoms"`
}

type AnteFamily struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Disease []string `json:"disease"`
}

type AnteFamilyOutput struct {
	ID      *string   `json:"id"`
	Name    *string   `json:"name"`
	Disease *[]string `json:"disease"`
}

type AnteFamilyInput struct {
	Name    string   `json:"name"`
	Disease []string `json:"disease"`
}

type Treatment struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Disease     string   `json:"disease"`
	Symptoms    []string `json:"symptoms"`
	SideEffects []string `json:"side_effects"`
}

type TreatmentOutput struct {
	ID          *string   `json:"id"`
	Name        *string   `json:"name"`
	Disease     *string   `json:"disease"`
	Symptoms    *[]string `json:"symptoms"`
	SideEffects *[]string `json:"side_effects"`
}

type TreatmentInput struct {
	Name        string   `json:"name"`
	Disease     string   `json:"disease"`
	Symptoms    []string `json:"symptoms"`
	SideEffects []string `json:"side_effects"`
}

/**************** GraphQL types *****************/

type DiseasesResponse struct {
	GetDiseases []Disease `json:"getDiseases"`
}

type getAlertByIdResponse struct {
	Content AlertOutput `json:"getAlertByIdResponse"`
}

type getAlertsResponse struct {
	GetAlerts []Alert `json:"getAlerts"`
}

type createAlertResponse struct {
	Content AlertOutput `json:"createAlertResponse"`
}

type getDiseaseByIdResponse struct {
	Content DiseaseOutput `json:"getDiseaseByIdResponse"`
}

type createDiseaseResponse struct {
	Content DiseaseOutput `json:"createDiseaseResponse"`
}

type getSymptomByIdResponse struct {
	Content SymptomOutput `json:"getSymptomByIdResponse"`
}

type createSymptomeResponse struct {
	Content SymptomOutput `json:"createSymptomeResponse"`
}

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

type getTreatmentByIdResponse struct {
	Content TreatmentOutput `json:"getTreatmentById"`
}

type createTreatmentResponse struct {
	Content TreatmentOutput `json:"createTreatment"`
}

/*************** Implementations *****************/

func GetAlerts() ([]Alert, error) {
	var resp getAlertsResponse
	var err error

	query := `query getAlerts {
                getAlertById {
                    id,
					name,
                    sex,
					height,
                    weight,
					symptoms,
					comment,
                }
            }`

	err = Query(query, nil, &resp)
	return resp.GetAlerts, err
}

func GetAlertsHotFix() []Alert {
	var grossesse Alert

	grossesse.ID = "65a37163854654b936c39ce2"
	grossesse.Name = "grossesse_extra_uterine"
	grossesse.Sex = new(string)
	*grossesse.Sex = "F"
	grossesse.Symptoms = []string{"abdominalgies"}
	grossesse.Comment = "Toute femme en âge de procréer ayant des douleurs abdominales ou des saignements est suspecte de grossesse extra-utérine (GEU) jusqu'à preuve du contraire"

	var etatgeneral Alert
	etatgeneral.ID = "65afbcc162c5259b6f30cc4e"
	etatgeneral.Name = "alteration_etat_general"
	etatgeneral.Symptoms = []string{"asthenie", "anorexie", "amaigrissement"}
	etatgeneral.Comment = "L'asthénie, l'anorexie et l'amaigrissement constituent le triptyque de l'altération de l'état général. Il peut être nécessaire de faire des tests concerant un probable cancer"

	return []Alert{grossesse, etatgeneral}
}

func getAlertById(id string) (Alert, error) {
	var alert getAlertByIdResponse
	var resp Alert
	query := `query getAlertById($id: String!) {
                getAlertById(id: $id) {
                    id,
					name,
                    sex,
					height,
                    weight,
					symptoms,
					comment,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &alert)
	_ = copier.Copy(&resp, &alert.Content)
	return resp, err
}

func createAlert(newDisease AlertInput) (Alert, error) {
	var alert createAlertResponse
	var resp Alert
	query := `mutation createDisease($name: String!, $sex: String!, $height: Int, $weight: Int, $symptoms: [String!]!, $comment: String!) {
        createDisease(name: $name, sex: $sex, height: $height, weight: $weight symptoms: $symptoms, comment: $comment) {
                    id,
					name,
                    sex,
					height,
                    weight,
					symptoms,
					comment,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":     newDisease.Name,
		"sex":      newDisease.Sex,
		"height":   newDisease.Height,
		"weight":   newDisease.Weight,
		"symptoms": newDisease.Symptoms,
		"comment":  newDisease.Comment,
	}, &alert)
	_ = copier.Copy(&resp, &alert.Content)
	return resp, err
}

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

func getPossibleSymptoms() []string {
	return []string{"respiration_difficile", "toux", "respiration_sifflante", "somnolence", "anxiete", "brulure_poitrine", "respiration_difficile", "boule_gorge", "maux_de_tetes", "vision_trouble", "tache_visuel", "abdominalgies", "asthenie", "anorexie", "amaigrissement"}
}

func getDiseaseById(id string) (Disease, error) {
	var disease getDiseaseByIdResponse
	var resp Disease
	query := `query getDiseaseById($id: String!) {
                getDiseaseById(id: $id) {
                    id,
					code,
                    name,
					symptoms,
					symptoms_acute,
                    symptoms_subacute,
					symptoms_chronic,
					advice,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &disease)
	_ = copier.Copy(&resp, &disease.Content)
	return resp, err
}

func createDisease(newDisease DiseaseInput) (Disease, error) {
	var disease createDiseaseResponse
	var resp Disease
	query := `mutation createDisease($code: String!, $name: String!, $symptoms: [string!]! $symptoms_acute: Map, $symptoms_subacute: Map, $symptoms_chronic: Map, $advice: String) {
        createDisease(code: $code, name: $name, symptoms: $symptoms, symptoms_acute: $symptoms_acute, symptoms_subacute: $symptoms_subacute, symptoms_chronic: $symptoms_chronic, advice: $advice) {
                    id,
					code,
                    name,
					symptoms,
					symptoms_acute,
                    symptoms_subacute,
					symptoms_chronic,
					advice,
                }
            }`
	err := Query(query, map[string]interface{}{
		"code":              newDisease.Code,
		"name":              newDisease.Name,
		"symptoms":          newDisease.Symptoms,
		"symptoms_acute":    newDisease.SymptomsAcute,
		"symptoms_subacute": newDisease.SymptomsSubacute,
		"symptoms_chronic":  newDisease.SymptomsChronic,
		"advice":            newDisease.Advice,
	}, &disease)
	_ = copier.Copy(&resp, &disease.Content)
	return resp, err
}

func getSymptomById(id string) (Symptom, error) {
	var symptom getSymptomByIdResponse
	var resp Symptom
	query := `query getSymptomById($id: String!) {
                getSymptomById(id: $id) {
                    id,
					code,
                    name,
					location,
                    duration,
					acute,
					subacute,
					chronic,
					advice,
					question,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &symptom)
	_ = copier.Copy(&resp, &symptom.Content)
	return resp, err
}

func createSymptom(newSymptom SymptomInput) (Symptom, error) {
	var symptom createSymptomeResponse
	var resp Symptom
	query := `mutation createSymptom($code: String!, $name: String!, $location: String, $duration: Int, $acute: Int, $subacute: Int, $chronic: Int, $advice: String, $question: String!) {
        createSymptom(code: $code, name: $name, location: $location, duration: $duration, acute: $acute, subacute: $subacute, chronic: $chronic, advice: $advice, question: $question) {
                    id,
					code,
                    name,
					location,
                    duration,
					acute,
					subacute,
					chronic,
					advice,
					question,
                }
            }`
	err := Query(query, map[string]interface{}{
		"code":     newSymptom.Code,
		"name":     newSymptom.Name,
		"location": newSymptom.Location,
		"duration": newSymptom.Duration,
		"acute":    newSymptom.Acute,
		"subacute": newSymptom.Subacute,
		"chronic":  newSymptom.Chronic,
		"advice":   newSymptom.Advice,
		"question": newSymptom.Question,
	}, &symptom)
	_ = copier.Copy(&resp, &symptom.Content)
	return resp, err
}

func getAnteChirById(id string) (AnteChir, error) {
	var antechir getAnteChirByIdResponse
	var resp AnteChir
	query := `query getAnteChirById($id: String!) {
                getAnteChirById(id: $id) {
                    id,
                    name,
                    localisation,
					induced_symptoms,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &antechir)
	_ = copier.Copy(&resp, &antechir.Content)
	return resp, err
}

func createAnteChir(newAnteChir AnteChirInput) (AnteChir, error) {
	var antechir createAnteChirResponse
	var resp AnteChir
	query := `mutation createAnteChir($name: String!, $localisation: String!, $induced_symptoms: [SymptomInput!]) {
        createAnteChir(name:$name, localisation:$localisation, induced_symptoms:$induced_symptoms) {
                    id,
					name,
                    localisation,
                    induced_symptoms,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":             newAnteChir.Name,
		"localisation":     newAnteChir.Localisation,
		"induced_symptoms": newAnteChir.InducedSymptoms,
	}, &antechir)
	_ = copier.Copy(&resp, &antechir.Content)
	return resp, err
}

func getAnteDiseaseById(id string) (AnteDisease, error) {
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

func createAnteDisease(newAnteDisease AnteDiseaseInput) (AnteDisease, error) {
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

func getAnteFamilyById(id string) (AnteFamily, error) {
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

func createAnteFamily(newAnteFamily AnteFamilyInput) (AnteFamily, error) {
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

func getTreatmentById(id string) (Treatment, error) {
	var treatment getTreatmentByIdResponse
	var resp Treatment
	query := `query getTreatmentById($id: String!) {
                getTreatmentById(id: $id) {
                    id,
                    name,
                    disease,
					symptoms,
					side_effects,
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &treatment)
	_ = copier.Copy(&resp, &treatment.Content)
	return resp, err
}

func createTreatment(newTreatment TreatmentInput) (Treatment, error) {
	var treatment createTreatmentResponse
	var resp Treatment
	query := `mutation createTreatment($name: String!, $disease: DiseaseInput!, $symptoms: [SymptomInput!]!, $side_effects: [SymptomInput]) {
        createTreatment(name: $name, disease: $disease, symptoms: $symptoms, side_effects: $side_effects) {
                    id,
                    name,
                    disease,
					symptoms,
					side_effects,
                }
            }`
	err := Query(query, map[string]interface{}{
		"name":         newTreatment.Name,
		"disease":      newTreatment.Disease,
		"symptoms":     newTreatment.Symptoms,
		"side_effects": newTreatment.SideEffects,
	}, &treatment)
	_ = copier.Copy(&resp, &treatment.Content)
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
