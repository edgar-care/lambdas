package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/chat"
	"github.com/edgar-care/edgarlib/v2/double_auth"
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

	accountID := lib.AuthMiddleware(input.Payload.AuthToken)
	check_account := authlib.CheckAccountEnable(accountID)
	if check_account.Code == 409 || check_account.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": check_account.Err.Error(),
		}, check_account.Code)
		return
	}

	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	createdChat := edgarlib.CreateChat(accountID, edgarlib.ContentInput{Message: input.Payload.Message, RecipientIds: input.Payload.RecipientIds})
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

	var deviceIds []string
	for _, recipientID := range input.Payload.RecipientIds {
		connectedDevice := double_auth.GetDeviceConnect(recipientID, 0, 0)
		if connectedDevice.Err != nil {
			continue
		}
		for _, connectedDevice := range connectedDevice.DevicesConnect {
			deviceIds = append(deviceIds, connectedDevice.ID)
		}
	}

	lib.BroadcastMessage(w, deviceIds, data)

	response := map[string]interface{}{
		"message": "Message sent",
	}

	lib.WriteResponse(w, response, 200)
}
