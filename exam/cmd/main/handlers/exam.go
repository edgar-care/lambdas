package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/exam/cmd/main/lib"
	"github.com/edgar-care/exam/cmd/main/services"
)

type examInput struct {
	Context []services.ExamContextItem `json:"context"`
}

func Exam(w http.ResponseWriter, req *http.Request) {
	var input examInput
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	question, possibleSymptoms, isDone := services.GuessQuestion(input.Context)

	lib.WriteResponse(w, map[string]interface{}{
		"context":  input.Context,
		"question": question,
		"symptoms": possibleSymptoms,
		"done":     isDone,
	}, 200)
}
