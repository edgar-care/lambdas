package handlers

import (
	"encoding/json"
	lib "github.com/edgar-care/chat/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/chat"
	"net/http"
)

type ReadMessageInput struct {
	ConnectionId string             `json:"connectionId"`
	Payload      PayloadSendMessage `json:"payload"`
}

type PayloadReadMessage struct {
	AuthToken string `json:"authToken"`
	ChatId    string `json:"chatId"`
}

func ReadMessage(w http.ResponseWriter, req *http.Request) {

	var input ReadMessageInput

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

	readMessage := edgarlib.UpdateMessageRead(id, input.Payload.ChatId)
	if readMessage.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": readMessage.Err.Error(),
		}, readMessage.Code)
		return
	}

	data := map[string]interface{}{
		"action": "read_message",
		"chat":   readMessage.Chat,
	}

	lib.WriteResponse(w, data, 200)

}
