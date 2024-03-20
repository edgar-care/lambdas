package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib"
	edgar_diag "github.com/edgar-care/edgarlib/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/http"
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

	resp := edgar_diag.Diagnose(input.Id, input.Sentence)

	if resp.Err != nil {
		edgarhttp.WriteResponse(w, map[string]interface{}{
			"message": resp.Err.Error(),
		}, resp.Code)
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"done":     resp.Done,
		"question": resp.Question,
	}, 200)
}
