package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
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
