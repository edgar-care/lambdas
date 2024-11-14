package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/chat"
	"github.com/edgar-care/edgarlib/v2/double_auth"
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

	sendMessage := edgarlib.AddMessageChat(accountID, edgarlib.ContentMessage{Message: input.Payload.Message, ChatId: input.Payload.ChatId})
	if sendMessage.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": sendMessage.Err.Error(),
		}, sendMessage.Code)
		return
	}

	data := map[string]interface{}{
		"action":    "receive_message",
		"owner_id":  accountID,
		"message":   input.Payload.Message,
		"timestamp": sendMessage.Chat.Messages[len(sendMessage.Chat.Messages)-1].SendedTime,
		"chat_id":   sendMessage.Chat.ID,
	}

	var deviceIds []string
	for _, participants := range sendMessage.Chat.Participants {
		connectedDevice := double_auth.GetDeviceConnect(participants.ParticipantID, 0, 0)
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
