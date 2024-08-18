package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/chat"
	"net/http"
)

type SendMessageInput struct {
	ConnectionId string             `json:"connectionId"`
	Payload      PayloadSendMessage `json:"payload"`
}

type PayloadSendMessage struct {
	AuthToken string `json:"authToken"`
	Message   string `json:"message"`
	ChatId    string `json:"chat_id"`
}

func SendMessage(w http.ResponseWriter, req *http.Request) {

	var input SendMessageInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	DoctorID := lib.AuthMiddlewareDoctor(input.Payload.AuthToken)
	PatientID := lib.AuthMiddleware(input.Payload.AuthToken)
	if DoctorID == "" && PatientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	id := DoctorID + PatientID

	sendMessage := edgarlib.AddMessageChat(id, edgarlib.ContentMessage{Message: input.Payload.Message, ChatId: input.Payload.ChatId})
	if sendMessage.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": sendMessage.Err.Error(),
		}, sendMessage.Code)
		return
	}

	data := map[string]interface{}{
		"action":    "receive_message",
		"owner_id":  id,
		"message":   input.Payload.Message,
		"timestamp": sendMessage.Chat.Messages[len(sendMessage.Chat.Messages)-1].SendedTime,
		"chat_id":   sendMessage.Chat.ID,
	}

	var recipientIds []string
	for _, participants := range sendMessage.Chat.Participants {
		recipientIds = append(recipientIds, participants.ParticipantID)
	}

	lib.BroadcastMessage(w, recipientIds, data)

	response := map[string]interface{}{
		"message": "Message sent",
	}

	lib.WriteResponse(w, response, 200)
}
