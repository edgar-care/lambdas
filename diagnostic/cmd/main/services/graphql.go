package services

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
	"os"
)

/********** Types ***********/

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

type Session struct {
	Id           string   `json:"id"`
	Symptoms     []string `json:"symptoms"`
	Age          int      `json:"age"`
	Height       int      `json:"height"`
	Weight       int      `json:"weight"`
	Sex          string   `json:"sex"`
	LastQuestion string   `json:"last_question"`
	Logs         []Logs   `json:"logs"`
	Alerts       []string `json:"alerts"`
}

type SessionOutput struct {
	Id           string    `json:"id"`
	Symptoms     *[]string `json:"symptoms"`
	Age          *int      `json:"age"`
	Height       *int      `json:"height"`
	Weight       *int      `json:"weight"`
	Sex          *string   `json:"sex"`
	LastQuestion *string   `json:"last_question"`
	Logs         *[]Logs   `json:"logs"`
	Alerts       *[]string `json:"alerts"`
}

type SessionInput struct {
	Symptoms     []string `json:"symptoms"`
	Age          int      `json:"age"`
	Height       int      `json:"height"`
	Weight       int      `json:"weight"`
	Sex          string   `json:"sex"`
	LastQuestion string   `json:"last_question,omitempty"`
	Logs         []Logs   `json:"logs"`
	Alerts       []string `json:"alerts"`
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
	query := `mutation createSession($symptoms: [String!]!, $age: Int!, $height: Int!, $weight: Int!, $sex: String!, $last_question: String!, $logs: [LogsInput!]!, $alerts: [String!]!) {
            createSession(symptoms:$symptoms, age:$age, height:$height, weight:$weight, sex:$sex, last_question:$last_question, logs: $logs, alerts: $alerts) {
                    id,
					symptoms,
					age,
					height,
					weight,
					sex,
					last_question,
					logs {
						question,
						answer
					},
					alerts
                }
            }`
	err := Query(query, map[string]interface{}{
		"symptoms":      newSession.Symptoms,
		"age":           newSession.Age,
		"height":        newSession.Height,
		"weight":        newSession.Weight,
		"sex":           newSession.Sex,
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
	query := `query getSessionById($id: String!) {
                getSessionById(id: $id) {
                    id,
					symptoms,
					age,
					height,
					weight,
					sex,
					last_question,
					logs {
						question,
						answer
					},
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
	query := `mutation updateSession($id: String!, $symptoms: [String!], $age: Int, $height: Int, $weight: Int, $sex: String, $last_question: String, $logs: [LogsInput!], $alerts: [String!]) {
                updateSession(id: $id, symptoms:$symptoms, age:$age, height:$height, weight:$weight, sex:$sex, last_question:$last_question, logs: $logs, alerts: $alerts) {
                    id,
					symptoms,
					age,
					height,
					weight,
					sex,
					last_question,
					logs {
						question,
						answer
					},
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
