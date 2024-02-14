package handlers

import (
	"net/http"
	"strings"

	"github.com/edgar-care/diagnostic/cmd/main/services"
	edgarhttp "github.com/edgar-care/edgarlib/http"
)

func Initiate(w http.ResponseWriter, req *http.Request) {
	var input services.SessionInput
	//err := json.NewDecoder(req.Body).Decode(&input)
	// QUICK FIX
	// TODO: Fix the model
	input.Symptoms = []services.SessionSymptom{}
	input.Age = 0
	input.Height = 0
	input.Weight = 0
	input.Sex = "M"
	input.LastQuestion = ""
	input.Logs = []services.Logs{}
	input.Alerts = []string{}

	services.WakeNlpUp()

	session, err := services.CreateSession(input)
	if err != nil {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Unable to create session: " + strings.ToLower(err.Error()[9:]),
		}, 400)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"sessionId": session.Id,
	}, 200)
}
