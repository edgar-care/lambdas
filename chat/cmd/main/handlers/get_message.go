package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/chat"
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

	allMessage := edgarlib.GetChat(accountID)
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
