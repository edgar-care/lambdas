package handlers

import (
	"net/http"
	"strings"

	"github.com/edgar-care/diagnostic/cmd/main/services"
	edgarhttp "github.com/edgar-care/edgarlib/http"
)

func getPatient(patientId string) services.SessionInput {
	var newSession services.SessionInput

	//patient := getPatientById(patientId)
	//patient_infos := getInfoById(patient.OnboardingInfoID)
	//patient_health := getHealthByID(patient.OnboardingHealthID)

	// TODO: replace with datas from patient_info and patient_health
	newSession.Symptoms = []services.SessionSymptom{}
	newSession.Age = 25
	newSession.Height = 180
	newSession.Weight = 75
	newSession.Sex = "M"
	newSession.AnteChirs = []string{}
	newSession.AnteDiseases = []string{}
	newSession.Treatments = []string{}
	newSession.LastQuestion = ""
	newSession.Logs = []services.Logs{}
	newSession.Alerts = []string{}

	return newSession
}

func Initiate(w http.ResponseWriter, req *http.Request) {
	//var input services.SessionInput
	input := getPatient("patientId")
	//err := json.NewDecoder(req.Body).Decode(&input)
	// QUICK FIX
	// TODO: Fix the model

	//input.Symptoms = []services.SessionSymptom{}
	//input.Age = 25
	//input.Height = 180
	//input.Weight = 75
	//input.Sex = "M"
	//input.AnteChirs = []string{}
	//input.AnteDiseases = []string{}
	//input.Treatments = []string{}
	//input.LastQuestion = ""
	//input.Logs = []services.Logs{}
	//input.Alerts = []string{}

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
