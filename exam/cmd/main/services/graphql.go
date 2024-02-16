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

type getSymptomsResponse struct {
	GetSymptoms []Symptom `json:"getDiseases"`
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
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
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
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
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
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
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
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
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
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff052ab205810761ccac8",
				"code": "angine",
				"name": "angine",
				"symptoms": [
					"pharyngodynie",
					"dysphagie",
					"pharyngite",
					"adenopathie",
					"pyrexie",
					"cephalee",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff0a1ab205810761ccac9",
				"code": "rhinopharyngite_aigue",
				"name": "rhinopharyngite_aigue",
				"symptoms": [
					"rhinorrhee",
					"sternutation",
					"pharyngite",
					"pharyngodynie",
					"toux_non_productive",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff0f2ab205810761ccaca",
				"code": "gastro-enterite_aigue",
				"name": "gastro-enterite_aigue",
				"symptoms": [
					"diarrhee",
					"vomissements",
					"abdominalgies",
					"nausees",
					"pyrexie",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff19eab205810761ccacb",
				"code": "otite",
				"name": "otite",
				"symptoms": [
					"pyrexie",
					"irritabilite",
					"isomnie",
					"hypoacousie",
					"otorrhee"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff1eeab205810761ccacc",
				"code": "diabete",
				"name": "diabete",
				"symptoms": [
					"polyurie",
					"polydipsie",
					"polyphagie",
					"retard_de_cicatrisation",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff53ed5394493fa1872e5",
				"code": "hypercholesterolemie",
				"name": "hypercholesterolemie",
				"symptoms": [
					"xanthomes",
					"arcs_corneens",
					"douleurs_thoraciques",
					"myalgies",
					"troubles_de_la_circulation"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff580d5394493fa1872e6",
				"code": "hypertension_arterielle",
				"name": "hypertension_arterielle",
				"symptoms": [
					"cephalee",
					"vertiges",
					"etourdissements",
					"vision_trouble",
					"epistaxis",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff5d4d5394493fa1872e7",
				"code": "hypothyroidie",
				"name": "hypothyroidie",
				"symptoms": [
					"prise_de_poids",
					"frilosite",
					"xerose_cutanee",
					"constipation",
					"onychorhexie",
					"Bradycardie"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff623d5394493fa1872e8",
				"code": "hyperthyroidie",
				"name": "hyperthyroidie",
				"symptoms": [
					"amaigrissement",
					"palpitations",
					"nervosite",
					"irritabilite",
					"tremblements",
					"intolerance_a_la_chaleur",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
				"advice": null
			},
			{
				"id": "65aff662d5394493fa1872e9",
				"code": "burn_out",
				"name": "burn_out",
				"symptoms": [
					"epuisement_emotionnel",
					"detachement_emotionnel ",
					"insomnie",
					"irritabilite",
					"syndrome_imposteur",
					"somatisation",
					"fatigue"
				],
				"symptoms_acute": null,
				"symptoms_subacute": null,
				"symptoms_chronic": null,
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

func GetSymptoms() ([]Symptom, error) {
	var resp getSymptomsResponse
	var err error

	jsonData := `{
		"getSymptoms": [
			{
				"id": "64c66b9107d76b7bdd5d6774",
				"code": "vision_trouble",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"trouble",
					"vision"
				],
				"advice": "",
				"question": "Avez vous la vision trouble ?"
			},
			{
				"id": "64c66bbc07d76b7bdd5d6775",
				"code": "fievre",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"froid",
					"temperature",
					"fievre"
				],
				"advice": null,
				"question": "Avez vous de la fièvre ?"
			},
			{
				"id": "64c66be607d76b7bdd5d6776",
				"code": "maux_de_ventre",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"ventre",
					"estomac",
					"intestin"
				],
				"advice": null,
				"question": "Avez vous des maux de ventre ?"
			},
			{
				"id": "64c66c1507d76b7bdd5d6777",
				"code": "vomissements",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"vomir",
					"vomissement",
					"brassé",
					"vomi"
				],
				"advice": null,
				"question": "Avez vous des vomissements ?"
			},
			{
				"id": "653012f3045016d0cb65a6f4",
				"code": "maux_de_tete",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"maux de tete",
					"mal de tete",
					"tete"
				],
				"advice": "reposez vous",
				"question": "Avez vous des maux de têtes ?"
			},
			{
				"id": "6538f9fec6a11ccf4143f452",
				"code": "respiration_difficile",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"respiration difficile",
					"difficulté"
				],
				"advice": null,
				"question": "Avez vous des difficultés pour respirer ?"
			},
			{
				"id": "6538fa1bc6a11ccf4143f453",
				"code": "toux",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"toux",
					"tousser"
				],
				"advice": null,
				"question": "Avez vous de la toux ?"
			},
			{
				"id": "6538fa9fc6a11ccf4143f454",
				"code": "respiration_sifflante",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"sifflante",
					"tousser"
				],
				"advice": null,
				"question": "Avez vous la respiration sifflante ?"
			},
			{
				"id": "6538fb72c6a11ccf4143f455",
				"code": "somnolence",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"somnolence",
					"envie de dormir"
				],
				"advice": null,
				"question": "Ressentez vous parfois soudainement l'envie de dormir ?"
			},
			{
				"id": "6538fb91c6a11ccf4143f456",
				"code": "anxiete",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"anxiete",
					"anxieux"
				],
				"advice": null,
				"question": "Êtes vous anxieux ?"
			},
			{
				"id": "6538fbe7c6a11ccf4143f457",
				"code": "brulure_poitrine",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"brulure dans la poitrine",
					"poitrine"
				],
				"advice": null,
				"question": "Ressentez vous des brulures dans la poitrine ?"
			},
			{
				"id": "65390c5bd792cad35642a098",
				"code": "boule_gorge",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"boule dans la gorge"
				],
				"advice": null,
				"question": "Ressentez vous des brulures dans la poitrine ?"
			},
			{
				"id": "65392842467ff3023b4631ec",
				"code": "tache_visuel",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"tache",
					"zone sombre",
					"vision"
				],
				"advice": null,
				"question": "Voyez vous comme des taches sombre ?"
			},
			{
				"id": "6539285c467ff3023b4631ed",
				"code": "vertige",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"vertige"
				],
				"advice": null,
				"question": "Avez vous des vertiges ?"
			},
			{
				"id": "65392887467ff3023b4631ee",
				"code": "perte_equilibre",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"equilibre",
					"perte d'équilibre"
				],
				"advice": null,
				"question": "Est-ce qu'il vous arrive de perdre l'équilibre ?"
			},
			{
				"id": "653928b60ef9f8b48f6e0a79",
				"code": "douleur_pulsatile",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"pulsion",
					"battement",
					"pulsatile"
				],
				"advice": null,
				"question": "Ressentez vous une douleur pulsatile ?"
			},
			{
				"id": "6539292f467ff3023b4631f0",
				"code": "sensibilite_lumiere",
				"name": "",
				"location": null,
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"lumiere",
					"sensible"
				],
				"advice": null,
				"question": "Êtes vous sensible à la lumiere ?"
			},
			{
				"id": "65af7b5885441f20c96fa2a8",
				"code": "pharyngodynie",
				"name": "pharyngodynie",
				"location": "cou",
				"duration": null,
				"acute": 0,
				"subacute": 7,
				"chronic": 28,
				"symptom": [
					"mal de gorge",
					"gorge",
					"gorge sèche"
				],
				"advice": null,
				"question": ""
			},
			{
				"id": "65af7cf085441f20c96fa2a9",
				"code": "dysphagie",
				"name": "dysphagie",
				"location": "cou",
				"duration": null,
				"acute": 0,
				"subacute": 14,
				"chronic": 84,
				"symptom": [
					"difficulté à avaler",
					"blocage dans la gorge",
					"mal quand avale"
				],
				"advice": null,
				"question": "Avez-vous des difficultés à avaler"
			},
			{
				"id": "65af7da685441f20c96fa2aa",
				"code": "pharyngite",
				"name": "pharyngite",
				"location": "cou",
				"duration": null,
				"acute": 0,
				"subacute": 7,
				"chronic": 14,
				"symptom": [
					"inflammation de la gorge",
					"irritation dans la gorge",
					"gorge qui brûle"
				],
				"advice": null,
				"question": "Avez-vous la gorge qui vous brûle ?"
			},
			{
				"id": "65af7efb85441f20c96fa2ab",
				"code": "adenopathie",
				"name": "adenopathie",
				"location": "",
				"duration": null,
				"acute": 0,
				"subacute": 28,
				"chronic": 84,
				"symptom": [
					"gonflement au cou",
					"gonflement à l'abdomene",
					"gonflement du thorax"
				],
				"advice": null,
				"question": "Sentez-vous un gonflement là où vous avez mal ?"
			},
			{
				"id": "65af7f3385441f20c96fa2ac",
				"code": "pyrexie",
				"name": "pyrexie",
				"location": "",
				"duration": null,
				"acute": 0,
				"subacute": 7,
				"chronic": 14,
				"symptom": [
					"fièvre"
				],
				"advice": null,
				"question": "Avez-vous de la fièvre ?"
			},
			{
				"id": "65af7f9785441f20c96fa2ad",
				"code": "cephalee",
				"name": "cephalee",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": 3,
				"symptom": [
					"maux de tête",
					"mal de tête",
					"migraine"
				],
				"advice": null,
				"question": "Avez-vous mal à la tête ?"
			},
			{
				"id": "65af80fc85441f20c96fa2ae",
				"code": "rhinorrhee",
				"name": "rhinorrhee",
				"location": "nez",
				"duration": null,
				"acute": 0,
				"subacute": null,
				"chronic": 20,
				"symptom": [
					"nez qui coule",
					"écoulement nasal"
				],
				"advice": null,
				"question": "Avez-vous le nez qui coule ?"
			},
			{
				"id": "65af818785441f20c96fa2af",
				"code": "sternutation",
				"name": "sternutation",
				"location": "nez",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"éternuements"
				],
				"advice": null,
				"question": "Éternuez-vous ?"
			},
			{
				"id": "65af81fa85441f20c96fa2b0",
				"code": "toux_non_productive",
				"name": "toux_non_productive",
				"location": "cou",
				"duration": null,
				"acute": 0,
				"subacute": 21,
				"chronic": 56,
				"symptom": [
					"toux sèche"
				],
				"advice": null,
				"question": "Avez-vous de la toux sèche ?"
			},
			{
				"id": "65af827485441f20c96fa2b1",
				"code": "diarrhée",
				"name": "diarrhée",
				"location": "",
				"duration": null,
				"acute": 0,
				"subacute": 7,
				"chronic": 14,
				"symptom": [
					"diarhée"
				],
				"advice": null,
				"question": "Avez-vous la diarrhée ?"
			},
			{
				"id": "65af829e85441f20c96fa2b2",
				"code": "vomissements",
				"name": "vomissements",
				"location": "",
				"duration": null,
				"acute": 0,
				"subacute": null,
				"chronic": 4,
				"symptom": [
					"vomissements"
				],
				"advice": null,
				"question": "Avez-vous des vomissements?"
			},
			{
				"id": "65af82fc85441f20c96fa2b3",
				"code": "abdominalgies",
				"name": "abdominalgies",
				"location": "abdominaux",
				"duration": null,
				"acute": 0,
				"subacute": 7,
				"chronic": 21,
				"symptom": [
					"douleurs abdominales"
				],
				"advice": null,
				"question": "Avez-vous des douleurs abdominales ?"
			},
			{
				"id": "65af839e85441f20c96fa2b4",
				"code": "nausees",
				"name": "nausees",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": 4,
				"symptom": [
					"nausées",
					"envie de vomir",
					"nauséeux"
				],
				"advice": null,
				"question": "Avez-vous des nausées ?"
			},
			{
				"id": "65af846c85441f20c96fa2b5",
				"code": "otalgie",
				"name": "otalgie",
				"location": "oreille",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": 7,
				"symptom": [
					"mal à l'oreille",
					"douleur à l'oreille"
				],
				"advice": null,
				"question": "Avez-vous mal à l'oreille ?"
			},
			{
				"id": "65af84ea85441f20c96fa2b6",
				"code": "troubles_de_l’equilibre",
				"name": "troubles_de_l’equilibre",
				"location": "oreille",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": 34,
				"symptom": [
					"problèmes d'équilibre",
					"troubles de l'équilibre"
				],
				"advice": null,
				"question": "Ressentez-vous des problèmes d'équilibre ?"
			},
			{
				"id": "65af855c85441f20c96fa2b7",
				"code": "inflammation",
				"name": "inflammation",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"inflammation",
					"rougeurs",
					"gonflements"
				],
				"advice": null,
				"question": "avez-vous des rougeurs ou des gonflement là où vous avez mal ?"
			},
			{
				"id": "65af860e85441f20c96fa2b8",
				"code": "polyurie",
				"name": "polyurie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"urines excessives",
					"beaucoup pipi"
				],
				"advice": null,
				"question": "est-ce que vous buvez et urinez beaucoup plus ue d'habitude ?"
			},
			{
				"id": "65af865d85441f20c96fa2b9",
				"code": "polydipsie",
				"name": "polydipsie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"soif excessives",
					"boire beaucoup"
				],
				"advice": null,
				"question": "est-ce que vous buvez beaucoup plus que d'habitude ?"
			},
			{
				"id": "65af86c385441f20c96fa2ba",
				"code": "polyphagie",
				"name": "polyphagie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"polyphagie",
					"hyperphagie",
					"faim excessive",
					"manger beaucoup"
				],
				"advice": null,
				"question": "Est-ce que vous souffrez d'une faim excessive ?"
			},
			{
				"id": "65af870685441f20c96fa2bb",
				"code": "vision_floue",
				"name": "vision_floue",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"vision floue",
					"vision trouble"
				],
				"advice": null,
				"question": "Est-ce que vous souffrez d'une vision trouble ?"
			},
			{
				"id": "65af876f85441f20c96fa2bc",
				"code": "retard_de_cicatrisation",
				"name": "retard_de_cicatrisation",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"retard de cicatrisation",
					"cicatrisation lente"
				],
				"advice": null,
				"question": "Est-ce que sentez que vous prenez beaucoup plus de temps à cicatriser ?"
			},
			{
				"id": "65af886485441f20c96fa2bd",
				"code": "xanthomes",
				"name": "xanthomes",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"xanthomes",
					"plaques jaune",
					"plaques jaunatre",
					"bosses"
				],
				"advice": null,
				"question": "Avez-vous remarqué des plaques jaunâtres ou des bosses sous la peau ?"
			},
			{
				"id": "65afc22262c5259b6f30cc4f",
				"code": "asthenie",
				"name": "asthenie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"asthenie",
					"sensation de fatigue",
					"sensation de faiblesse"
				],
				"advice": null,
				"question": "Avez-vous remarqué une sensation de faiblesse ou de fatigue sans changement dans votre mode de vie ou votre quotidien ?"
			},
			{
				"id": "65afc48462c5259b6f30cc50",
				"code": "anorexie",
				"name": "anorexie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"anorexie",
					"perte de l'envie de manger",
					"perte de l'appétit"
				],
				"advice": null,
				"question": "Avez-vous des rituels ou des comportements alimentaires spécifiques, tels que couper les aliments en petits morceaux, éviter certains groupes alimentaires, ou sauter fréquemment des repas ?"
			},
			{
				"id": "65afc54862c5259b6f30cc51",
				"code": "amaigrissement",
				"name": "amaigrissement",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"amaigrissement",
					"perte de poids"
				],
				"advice": null,
				"question": "Avez-vous remarqué une perte de poids significative récemment ?"
			},
			{
				"id": "65afe4e0ab205810761ccaaf",
				"code": "gerontoxon",
				"name": "gerontoxon",
				"location": "oeil",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"gerontoxon",
					"Arcs cornéens",
					"arc de la cornée"
				],
				"advice": null,
				"question": "Éprouvez-vous une gêne oculaire, une douleur ou une rougeur oculaire"
			},
			{
				"id": "65afe531ab205810761ccab0",
				"code": "douleurs_thoraciques",
				"name": "douleurs_thoraciques",
				"location": "thoraxe",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"douleurs thoraciques",
					"mal au torse",
					"mal au thoraxe"
				],
				"advice": null,
				"question": "Éprouvez-vous une douleur au niveau du thoraxe"
			},
			{
				"id": "65afe56fab205810761ccab1",
				"code": "myalgies",
				"name": "myalgies",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"myalgies",
					"faiblesse musculaire"
				],
				"advice": null,
				"question": "Éprouvez-vous une faiblesse musculaire ?"
			},
			{
				"id": "65afe608ab205810761ccab2",
				"code": "troubles_de_la_circulation",
				"name": "troubles_de_la_circulation",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"troubles de la circulation",
					"troubles de la circulation sanguine"
				],
				"advice": null,
				"question": "Avez-vous des engourdissements ou des picotements dans les membres ou des changements de couleur de la peau, tels que des zones pâles ou bleuâtres ?"
			},
			{
				"id": "65afe630ab205810761ccab3",
				"code": "vertiges",
				"name": "vertiges",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"vertiges"
				],
				"advice": null,
				"question": "Avez-vous des vertiges ?"
			},
			{
				"id": "65afe657ab205810761ccab4",
				"code": "étourdissement",
				"name": "étourdissement",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"étourdissement"
				],
				"advice": null,
				"question": "souffrez-vous d'étourdissements ?"
			},
			{
				"id": "65afe6c1ab205810761ccab5",
				"code": "dyspnée",
				"name": "dyspnée",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"dyspnée",
					"Essoufflement"
				],
				"advice": null,
				"question": "souffrez-vous d'essoufflements sans un effort physique particuler ?"
			},
			{
				"id": "65afe6efab205810761ccab6",
				"code": "epistaxis",
				"name": "epistaxis",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Épistaxis",
					"Saignement de nez"
				],
				"advice": null,
				"question": "saignez-vous souvent du nez ?"
			},
			{
				"id": "65afe737ab205810761ccab7",
				"code": "prise_de_poids",
				"name": "prise_de_poids",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"prise de poids"
				],
				"advice": null,
				"question": "avez-vous pris du poids soudainement et sans changement particulier dans votre mode de vie ou vos habitudes alimentaire ?"
			},
			{
				"id": "65afe7c1ab205810761ccab8",
				"code": "frilosite",
				"name": "frilosite",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"frilosite",
					"sensibilité accrue au froid"
				],
				"advice": null,
				"question": "ressentez-vous une sensibilité accrue au froid ?"
			},
			{
				"id": "65afe83cab205810761ccab9",
				"code": "xerose_cutanee",
				"name": "xerose_cutanee",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Xérose cutanée",
					"peau sèche"
				],
				"advice": null,
				"question": "Éprouvez-vous des démangeaisons des fissures de la peau  ?"
			},
			{
				"id": "65afe895ab205810761ccaba",
				"code": "constipation",
				"name": "constipation",
				"location": "",
				"duration": null,
				"acute": 0,
				"subacute": 3,
				"chronic": 21,
				"symptom": [
					"constipation"
				],
				"advice": null,
				"question": "souffrez-vous de constipations ?"
			},
			{
				"id": "65afe979ab205810761ccabb",
				"code": "onychorhexie",
				"name": "onychorhexie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"onychorhexie",
					"cheveux fragile",
					"ongles fragiles",
					"cheveux et ongles fragile"
				],
				"advice": null,
				"question": "Avez-vous les cheveux ou les ongles fragiles ou ressentez-vous des picotements ou des démangeaisons au niveau des ongles ?"
			},
			{
				"id": "65afeaadab205810761ccabc",
				"code": "bradycardie",
				"name": "bradycardie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"bradycardie",
					"ralentissement du rythme cardiaque"
				],
				"advice": null,
				"question": "souffrez-vous de bradycardie ?"
			},
			{
				"id": "65afead0ab205810761ccabd",
				"code": "troubles_menstruels",
				"name": "troubles_menstruels",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"troubles menstruels"
				],
				"advice": null,
				"question": "souffrez-vous de troubles menstruels ?"
			},
			{
				"id": "65afeb29ab205810761ccabe",
				"code": "palpitations",
				"name": "palpitations",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Palpitations"
				],
				"advice": null,
				"question": "souffrez-vous de palpitations ?"
			},
			{
				"id": "65afeb78ab205810761ccabf",
				"code": "nervosité",
				"name": "nervosité",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"nervosité"
				],
				"advice": null,
				"question": "Ressentez-vous fréquemment une sensation de nervosité ou d'anxiété ?"
			},
			{
				"id": "65afec13ab205810761ccac0",
				"code": "tremblements",
				"name": "tremblements",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"tremblements"
				],
				"advice": null,
				"question": "Ressentez-vous des tremblements ?"
			},
			{
				"id": "65afec97ab205810761ccac1",
				"code": "intolerance_a_la_chaleur",
				"name": "intolerance_a_la_chaleur",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Intolérance à la chaleur"
				],
				"advice": null,
				"question": "Ressentez-vous une intolérance à la chaleur ?"
			},
			{
				"id": "65afed10ab205810761ccac2",
				"code": "epuisement_emotionnel",
				"name": "epuisement_emotionnel",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Épuisement émotionnel"
				],
				"advice": null,
				"question": "Éprouvez-vous un épuisement émotionnel ?"
			},
			{
				"id": "65afed2aab205810761ccac3",
				"code": "detachement_emotionnel",
				"name": "detachement_emotionnel",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"Détachement émotionnel"
				],
				"advice": null,
				"question": "Éprouvez-vous un détachement émotionnel ?"
			},
			{
				"id": "65afed49ab205810761ccac4",
				"code": "insomnie",
				"name": "insomnie",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"insomnie"
				],
				"advice": null,
				"question": "souffrez-vous d'insomnies ?"
			},
			{
				"id": "65afeda1ab205810761ccac5",
				"code": "irritabilite",
				"name": "irritabilite",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"irritabilite"
				],
				"advice": null,
				"question": "vous sentez vous plus irritable ?"
			},
			{
				"id": "65afee26ab205810761ccac6",
				"code": "syndrome_imposteur",
				"name": "syndrome_imposteur",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"syndrome de l'imposteur"
				],
				"advice": null,
				"question": "Avez-vous des doutes quant à votre capacité à réussir ou à accomplir vos tâches professionnelles ?"
			},
			{
				"id": "65afeee0ab205810761ccac7",
				"code": "somatisation",
				"name": "somatisation",
				"location": "",
				"duration": null,
				"acute": null,
				"subacute": null,
				"chronic": null,
				"symptom": [
					"somatisation"
				],
				"advice": null,
				"question": "Ressentez-vous des douleurs physique accomagnant vos émotions négatives ?"
			}
		]
	}`

	if err = json.Unmarshal([]byte(jsonData), &resp); err != nil {
		fmt.Println("Error:", err)
	}
	return resp.GetSymptoms, err
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
