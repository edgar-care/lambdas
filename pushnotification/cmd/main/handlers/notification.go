package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"context"

	"github.com/edgar-care/pushnotification/cmd/main/lib"
	"github.com/edgar-care/pushnotification/cmd/main/services"


	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

//type PushNotificationRequest struct {
//	Token    string `json:"token"`
//	Title    string `json:"title"`
//	Message  string `json:"message"`
//}

func Notification(w http.ResponseWriter, r *http.Request) {
	// Initialize Firebase app
	var pushRequest services.NotificationInput
	opt := option.WithCredentialsFile("/home/tdarrieumerlou/egdar.care/lambdas/pushnotification/cmd/main/handlers/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Failed to create Firebase app: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Initialize Firebase messaging client
	client, err := app.Messaging(r.Context())
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Failed to create Firebase messaging client: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	// Parse request body
	//var pushRequest PushNotificationRequest
	err = json.NewDecoder(r.Body).Decode(&pushRequest)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// Create Firebase notification message
	message := &messaging.Message{
		Token: pushRequest.Token,
		Notification: &messaging.Notification{
			Title: pushRequest.Title,
			Body:  pushRequest.Message,
		},
	}

	// Send the notification
	response, err := client.Send(r.Context(), message)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Failed to send notification: " + err.Error(),
		}, http.StatusBadRequest)
		return
	}

	log.Printf("Successfully sent push notification. Response: %+v", response)

	lib.WriteResponse(w, map[string]interface{}{
		"notif": response,
	}, http.StatusOK)
}