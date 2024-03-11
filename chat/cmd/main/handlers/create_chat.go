package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/chat"
	"net/http"
)

type ChatInput struct {
	ConnectionId string            `json:"connectionId"`
	Payload      PayloadCreateChat `json:"payload"`
}

type PayloadCreateChat struct {
	AuthToken    string   `json:"authToken"`
	Message      string   `json:"message"`
	RecipientIds []string `json:"recipient_ids"`
}

func CreatChat(w http.ResponseWriter, req *http.Request) {
	var input ChatInput

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

	createdChat := edgarlib.CreateChat(id, edgarlib.ContentInput{Message: input.Payload.Message, RecipientIds: input.Payload.RecipientIds})
	if createdChat.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": createdChat.Err.Error(),
		}, createdChat.Code)
		return
	}

	data := map[string]interface{}{
		"action": "create_chat",
		"chat":   createdChat.Chat,
	}

	lib.BroadcastMessage(w, input.Payload.RecipientIds, data)

	response := map[string]interface{}{
		"message": "Message sent",
	}

	lib.WriteResponse(w, response, 200)
}
