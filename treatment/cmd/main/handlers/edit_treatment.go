package handlers

import (
	"encoding/json"
	edgarlib "github.com/edgar-care/edgarlib/treatment"
	"github.com/edgar-care/treatment/cmd/main/lib"
	"net/http"
)

func EditTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
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
