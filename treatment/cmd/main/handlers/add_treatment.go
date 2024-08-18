package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	edgarlib "github.com/edgar-care/edgarlib/v2/treatment"
	"github.com/edgar-care/treatment/cmd/main/lib"
)

func Addtreatment(w http.ResponseWriter, req *http.Request) {

	patientID := authlib.AuthMiddlewarePatient(w, req)
	if patientID.Code == 409 || patientID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": patientID.Err.Error(),
		}, patientID.Code)
		return
	}
	if patientID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateNewTreatmentInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	treatment := edgarlib.CreateTreatment(input, patientID.ID)

	if treatment.Err != nil {
		lib.WriteError(w, treatment.Code, treatment.Err.Error())
		return
	}

	response := map[string]interface{}{
		"treatment": map[string]interface{}{
			"name":           treatment.AnteDisease.Name,
			"still_relevant": treatment.AnteDisease.StillRelevant,
			"treatment":      treatment.Treatment,
		},
	}

	lib.WriteResponse(w, response, treatment.Code)
}
