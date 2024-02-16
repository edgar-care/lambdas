package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/edgar-care/diagnostic/cmd/main/services"
	"github.com/edgar-care/edgarlib"
	edgarhttp "github.com/edgar-care/edgarlib/http"
)

type diagnoseInput struct {
	Id       string
	Sentence string
}

func Diagnose(w http.ResponseWriter, req *http.Request) {
	var input diagnoseInput
	err := json.NewDecoder(req.Body).Decode(&input)
	edgarlib.CheckError(err)
	session, err := services.GetSessionById(input.Id)
	if err != nil {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Unable to get session: " + strings.ToLower(err.Error()[9:]),
		}, 400)
		return
	}

	//symptoms := services.StringToSymptoms(session.Symptoms)
	symptoms := session.Symptoms
	questionSymptom := []string{session.LastQuestion}
	if session.LastQuestion == "" {
		questionSymptom = []string{}

		tmp := services.Logs{Question: "", Answer: input.Sentence}
		session.Logs = append(session.Logs, tmp)
		edgarlib.CheckError(err)
	} else {
		tmp := services.Logs{Question: session.LastQuestion, Answer: input.Sentence}
		session.Logs = append(session.Logs, tmp)
		edgarlib.CheckError(err)
	}

	newSymptoms := services.CallNlp(input.Sentence, questionSymptom)
	for _, s := range newSymptoms.Context {
		symptoms = append(symptoms, s)
	}

	exam := services.CallExam(symptoms)
	if len(exam.Alert) > 0 {
		for _, alert := range exam.Alert {
			session.Alerts = append(session.Alerts, alert)
		}
	}
	//session.Symptoms = services.SymptomsToString(exam.Context)
	session.Symptoms = exam.Context
	if len(exam.Symptoms) > 0 {
		session.LastQuestion = exam.Symptoms[0]
	}
	fmt.Println(session.Id)
	_, err = services.UpdateSession(session)
	edgarlib.CheckError(err)

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"done":     exam.Done,
		"question": exam.Question,
	}, 200)
}
