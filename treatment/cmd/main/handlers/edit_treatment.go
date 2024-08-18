package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/treatment"
	"github.com/edgar-care/treatment/cmd/main/lib"
	"net/http"
)

func EditTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := authlib.AuthMiddlewarePatient(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	check_account := authlib.CheckAccountEnable(patientID)
	if check_account.Code == 409 {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authorized, this account is disable",
		}, 409)
		return
	}

	var input edgarlib.UpdateTreatmentInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	treatment := edgarlib.UpdateTreatment(input, patientID)

	if treatment.Err != nil {
		lib.WriteError(w, treatment.Code, treatment.Err.Error())
		return
	}

	//response := map[string]interface{}{
	//	"treatment": map[string]interface{}{
	//		"name":           treatment..Name,
	//		"still_relevant": treatment.Antedisease.StillRelevant,
	//		"treatment":      treatment.Treatment,
	//	},
	//}

	lib.WriteResponse(w, treatment.Treatment, treatment.Code)
}
