package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/chat"
	"net/http"
)

type GetMessageInput struct {
	ConnectionId string            `json:"connectionId"`
	Payload      PayloadGetMessage `json:"payload"`
}

type PayloadGetMessage struct {
	AuthToken string `json:"authToken"`
}

func GetMessages(w http.ResponseWriter, req *http.Request) {

	var input GetMessageInput

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

	allMessage := edgarlib.GetChat(id)
	if allMessage.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": allMessage.Err.Error(),
		}, allMessage.Code)
		return
	}

	data := map[string]interface{}{
		"action": "get_message",
		"chats":  allMessage.Chats,
	}

	lib.WriteResponse(w, data, 200)
}
