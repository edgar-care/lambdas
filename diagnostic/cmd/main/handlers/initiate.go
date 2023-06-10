package handlers

import (
	"encoding/json"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"github.com/edgar-care/diagnostic/cmd/main/services"
	"net/http"
	"strings"
)

func Initiate(w http.ResponseWriter, req *http.Request) {
	var input services.SessionInput
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)
	input.Symptoms = []string{}

	session, err := services.CreateSession(input)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to create session: " + strings.ToLower(err.Error()[9:]),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"sessionId": session.Id,
	}, 200)
}
