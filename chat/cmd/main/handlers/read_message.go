package handlers

import (
	"encoding/json"
	lib "github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/chat"
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

	accountID := lib.AuthMiddleware(input.Payload.AuthToken)
	check_account := authlib.CheckAccountEnable(accountID)
	if check_account.Code == 409 {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authorized, this account is disable",
		}, 409)
		return
	}

	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	readMessage := edgarlib.UpdateMessageRead(accountID, input.Payload.ChatId)
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
