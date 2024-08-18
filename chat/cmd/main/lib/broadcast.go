package lib

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2/redis"
	"net/http"
	"strings"
)

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
