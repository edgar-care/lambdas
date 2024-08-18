package handlers

import (
	"encoding/json"
	"net/http"

	edgarlib "github.com/edgar-care/edgarlib/follow_treatment"
	"github.com/edgar-care/treatment_follow_up/cmd/main/lib"
)

func AddFollowTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateNewFollowUpInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	follow_up := edgarlib.CreateTreatmentFollowUp(input, patientID)

	if follow_up.Err != nil {
		lib.WriteError(w, follow_up.Code, follow_up.Err.Error())
		return
	}

	lib.WriteResponse(w, follow_up.TreatmentFollowUp, follow_up.Code)
}
