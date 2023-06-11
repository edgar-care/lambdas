package handlers

import (
	"encoding/json"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"github.com/edgar-care/diagnostic/cmd/main/services"
	"net/http"
	"strings"
)

type diagnoseInput struct {
	Id       string
	Sentence string
}

func Diagnose(w http.ResponseWriter, req *http.Request) {
	var input diagnoseInput
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)
	session, err := services.GetSessionById(input.Id)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get session: " + strings.ToLower(err.Error()[9:]),
		}, 400)
		return
	}

	symptoms := services.StringToSymptoms(session.Symptoms)
	questionSymptom := []string{session.LastQuestion}
	if session.LastQuestion == "" {
		questionSymptom = []string{}
	}

	newSymptoms := services.CallNlp(input.Sentence, questionSymptom)
	for _, s := range newSymptoms.Context {
		symptoms = append(symptoms, s)
	}

	exam := services.CallExam(symptoms)
	session.Symptoms = services.SymptomsToString(exam.Context)
	if len(exam.Symptoms) > 0 {
		session.LastQuestion = exam.Symptoms[0]
	}
	_, err = services.UpdateSession(session)
	lib.CheckError(err)

	lib.WriteResponse(w, map[string]interface{}{
		"done":     exam.Done,
		"question": exam.Question,
	}, 200)
}
