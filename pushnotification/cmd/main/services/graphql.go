package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Notification struct {
	Id          string   `json:"id"`
	Token    string `json:"token"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

type NotificationOutput struct {
	Id           string    `json:"id"`
	Token    *string `json:"token"`
	Title    *string `json:"title"`
	Message  *string `json:"message"`
}

type NotificationInput struct {
	Token    string `json:"token"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}

/**************** GraphQL types *****************/

type createNotificationResponse struct {
	Content NotificationOutput `json:"createNotification"`
}

type getNotificationByIdResponse struct {
	Content NotificationOutput `json:"getNotificationById"`
}

/*************** Implementations *****************/

func CreateInfo(newNotification NotificationInput) (Notification, error) {
	var notification createNotificationResponse
	var resp Notification
	query := `mutation createInfo($token: String!, $title: String!, $message: String!) {
            createInfo(token:$token, title:$title, message:$message) {
                    id,
					token,
					title,
					message
                }
            }`
	err := Query(query, map[string]interface{}{
		"token":      newNotification.Token,
		"title":           newNotification.Title,
		"message":        newNotification.Message,
	}, &notification)
	_ = copier.Copy(&resp, &notification.Content)
	return resp, err
}

func GetInfoById(id string) (Notification, error) {
	var notification getNotificationByIdResponse
	var resp Notification
	query := `query getInfoById($id: String!) {
                getInfoById(id: $id) {
                    id,
					token,
					title,
					message
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &notification)
	_ = copier.Copy(&resp, &notification.Content)
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