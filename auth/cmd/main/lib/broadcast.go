package lib

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2/redis"
	"log"
	"net/http"
	"os"
	"strings"
)

func PostToConnection(w http.ResponseWriter, connectionId string, data []byte) {

	postBody, _ := json.Marshal(map[string]string{
		"connectionId": connectionId,
		"data":         string(data),
		"stage":        os.Getenv("STAGE"),
	})
	responseBody := bytes.NewBuffer(postBody)

	_, err := http.Post("http://x2025edgarcare2028075120000.francecentral.cloudapp.azure.com:8081/ws/send", "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
}

func DisconnectToConnection(w http.ResponseWriter, connectionId string) {

	disconnectBody, _ := json.Marshal(map[string]string{
		"connectionId": connectionId,
		"stage":        os.Getenv("STAGE"),
	})
	responseBody := bytes.NewBuffer(disconnectBody)

	_, err := http.Post("http://x2025edgarcare2028075120000.francecentral.cloudapp.azure.com:8081/ws/disconnect", "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}
}

func BroadcastMessage(w http.ResponseWriter, recipientIds []string, message map[string]interface{}) {

	for _, recipientID := range recipientIds {
		recipientConnectionId, err := redis.GetKey(recipientID)

		if err != nil || recipientConnectionId == "" {
			continue
		}

		jsonString, err := json.Marshal(message)
		if err != nil {
			WriteResponse(w, map[string]string{"message": err.Error()}, 500)
		}
		PostToConnection(w, strings.Replace(recipientConnectionId, "\n", "", -1), jsonString)

	}
}
