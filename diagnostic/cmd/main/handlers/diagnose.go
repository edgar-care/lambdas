package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2"
	edgarauth "github.com/edgar-care/edgarlib/v2/auth"
	edgar_diag "github.com/edgar-care/edgarlib/v2/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/v2/http"

	"net/http"
)

type diagnoseInput struct {
	Id       string `json:"id"`
	Sentence string `json:"sentence"`
}

func Diagnose(w http.ResponseWriter, req *http.Request) {
	var input diagnoseInput
	err := json.NewDecoder(req.Body).Decode(&input)
	edgarlib.CheckError(err)

	patientID := edgarauth.AuthMiddlewarePatient(w, req)
	if patientID == "" {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	resp := edgar_diag.Diagnose(input.Id, input.Sentence)

	if resp.Err != nil {
		edgarhttp.WriteResponse(w, map[string]interface{}{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"done":     resp.Done,
		"question": resp.Question,
	}, 200)
}
