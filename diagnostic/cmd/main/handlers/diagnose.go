package handlers

import (
	"encoding/json"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"github.com/edgar-care/edgarlib/v2"
	edgarauth "github.com/edgar-care/edgarlib/v2/auth"
	edgar_diag "github.com/edgar-care/edgarlib/v2/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/v2/http"

	"net/http"
)

type autoAnswerInput struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type diagnoseInput struct {
	Id         string           `json:"id"`
	Sentence   string           `json:"sentence"`
	AutoAnswer *autoAnswerInput `json:"auto_answer"`
}

func Diagnose(w http.ResponseWriter, req *http.Request) {
	var input diagnoseInput
	err := json.NewDecoder(req.Body).Decode(&input)
	edgarlib.CheckError(err)

	patientID := edgarauth.AuthMiddlewarePatient(w, req)
	if patientID.Code == 409 || patientID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": patientID.Err.Error(),
		}, patientID.Code)
		return
	}
	if patientID.ID == "" {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	resp := edgar_diag.Diagnose(input.Id, input.Sentence, (*edgar_diag.AutoAnswerinfo)(input.AutoAnswer))

	if resp.Err != nil {
		edgarhttp.WriteResponse(w, map[string]interface{}{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"done":        resp.Done,
		"question":    resp.Question,
		"auto_answer": resp.AutoAnswer,
	}, 200)
}
