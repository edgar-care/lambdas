package services

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
	"os"
)

/********** Sub-Structures ***********/

type SessionSymptom struct {
	Name     string `json:"name"`
	Presence *bool  `json:"presence"`
	Duration *int32 `json:"duration"`
}

type SessionSymptomInput struct {
	Name     string `json:"name"`
	Presence *bool  `json:"presence"`
	Duration *int32 `json:"duration"`
}

type SessionSymptomOutput struct {
	Name     string `json:"name"`
	Presence *bool  `json:"presence"`
	Duration *int32 `json:"duration"`
}

type Logs struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type LogsInput struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type LogsOutput struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

/********** Types ***********/

type Session struct {
	Id           string           `json:"id"`
	Symptoms     []SessionSymptom `json:"symptoms"`
	Age          int32            `json:"age"`
	Height       int32            `json:"height"`
	Weight       int32            `json:"weight"`
	Sex          string           `json:"sex"`
	AnteChirs    []string         `json:"ante_chirs"`
	AnteDiseases []string         `json:"ante_diseases"`
	Treatments   []string         `json:"treatments"`
	LastQuestion string           `json:"last_question"`
	Logs         []Logs           `json:"logs"`
	Alerts       []string         `json:"alerts"`
}

type SessionOutput struct {
	Id           string            `json:"id"`
	Symptoms     *[]SessionSymptom `json:"symptoms"`
	Age          *int32            `json:"age"`
	Height       *int32            `json:"height"`
	Weight       *int32            `json:"weight"`
	Sex          *string           `json:"sex"`
	AnteChirs    *[]string         `json:"ante_chirs"`
	AnteDiseases *[]string         `json:"ante_diseases"`
	Treatments   *[]string         `json:"treatments"`
	LastQuestion *string           `json:"last_question"`
	Logs         *[]Logs           `json:"logs"`
	Alerts       *[]string         `json:"alerts"`
}

type SessionInput struct {
	Symptoms     []SessionSymptom `json:"symptoms"`
	Age          int32            `json:"age"`
	Height       int32            `json:"height"`
	Weight       int32            `json:"weight"`
	Sex          string           `json:"sex"`
	AnteChirs    []string         `json:"ante_chirs"`
	AnteDiseases []string         `json:"ante_diseases"`
	Treatments   []string         `json:"treatments"`
	LastQuestion string           `json:"last_question"`
	Logs         []Logs           `json:"logs"`
	Alerts       []string         `json:"alerts"`
}

/**************** GraphQL types *****************/

type createSessionResponse struct {
	Content SessionOutput `json:"createSession"`
}

type getSessionByIdResponse struct {
	Content SessionOutput `json:"getSessionById"`
}

type updateSessionResponse struct {
	Content SessionOutput `json:"updateSessionResponse"`
}

/*************** Implementations *****************/

func CreateSession(newSession SessionInput) (Session, error) {
	var session createSessionResponse
	var resp Session
	query := `mutation createSession($symptoms: [SessionSymptomInput!]!, $age: Int!, $height: Int!, $weight: Int!, $sex: String!, $ante_chirs: [String!]!, $ante_diseases: [String!]!, $treatments: [String!]!, $last_question: String!, $logs: [LogsInput!]!, $alerts: [String!]!) {
				createSession(symptoms: $symptoms, age: $age, height: $height, weight: $weight, sex: $sex, ante_chirs: $ante_chirs, ante_diseases: $ante_diseases, treatments: $treatments, last_question: $last_question, logs: $logs, alerts: $alerts) {
					id
					symptoms {
						name
						presence
						duration
					}
					age
					height
					weight
					sex
					ante_chirs
					ante_diseases
					treatments
					last_question
					logs {
						question
						answer
					}
					alerts
				}
			}`
	err := Query(query, map[string]interface{}{
		"symptoms":      newSession.Symptoms,
		"age":           newSession.Age,
		"height":        newSession.Height,
		"weight":        newSession.Weight,
		"sex":           newSession.Sex,
		"ante_chirs":    newSession.AnteChirs,
		"ante_diseases": newSession.AnteDiseases,
		"treatments":    newSession.Treatments,
		"last_question": newSession.LastQuestion,
		"logs":          newSession.Logs,
		"alerts":        newSession.Alerts,
	}, &session)
	_ = copier.Copy(&resp, &session.Content)
	return resp, err
}

func GetSessionById(id string) (Session, error) {
	var session getSessionByIdResponse
	var resp Session
	query := `query	getSessionById($id: String!) {
				getSessionById(id: $id) {
					id
					symptoms {
						name
						presence
						duration
					}
					age
					height
					weight
					sex
					ante_chirs
					ante_diseases
					treatments
					last_question
					logs {
						question
						answer
					}
					alerts
				}
			}`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &session)
	_ = copier.Copy(&resp, &session.Content)
	return resp, err
}

func UpdateSession(newSession Session) (Session, error) {
	var session updateSessionResponse
	var resp Session
	query := `mutation updateSession($id: String!, $symptoms: [SessionSymptomInput!], $age: Int, $height: Int, $weight: Int, $sex: String, $ante_chirs: [String!], $ante_diseases: [String!], $treatments: [String!], $last_question: String, $logs: [LogsInput!], $alerts: [String!]) {
				updateSession(id: $id, symptoms: $symptoms, age: $age, height: $height, weight: $weight, sex: $sex, ante_chirs: $ante_chirs, ante_diseases: $ante_diseases, treatments: $treatments, last_question: $last_question, logs: $logs, alerts: $alerts) {
					id
					symptoms {
						name
						presence
						duration
					}
					age
					height
					weight
					sex
					ante_chirs
					ante_diseases
					treatments
					last_question
					logs {
						question
						answer
					}
					alerts
				}
			}`

	err := Query(query, map[string]interface{}{
		"id":            newSession.Id,
		"symptoms":      newSession.Symptoms,
		"age":           newSession.Age,
		"height":        newSession.Height,
		"weight":        newSession.Weight,
		"sex":           newSession.Sex,
		"ante_chirs":    newSession.AnteChirs,
		"ante_diseases": newSession.AnteDiseases,
		"treatments":    newSession.Treatments,
		"last_question": newSession.LastQuestion,
		"logs":          newSession.Logs,
		"alerts":        newSession.Alerts,
	}, &session)
	_ = copier.Copy(&resp, &session.Content)
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
