package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	edgarlib "github.com/edgar-care/edgarlib/follow_treatment"
	"github.com/edgar-care/treatment_follow_up/cmd/main/lib"
)

func GetFollowTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	follow_get := edgarlib.GetTreatmentFollowUpById(t)
	if follow_get.Err != nil {
		lib.WriteError(w, follow_get.Code, follow_get.Err.Error())
		return
	}

	lib.WriteResponse(w, follow_get.TreatmentFollowUp, follow_get.Code)
}

func GetfFollowsTreatments(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	follow_up := edgarlib.GetTreatmentFollowUp(patientID)
	if follow_up.Err != nil {
		lib.WriteError(w, follow_up.Code, follow_up.Err.Error())
		return
	}

	lib.WriteResponse(w, follow_up.TreatmentFollowUps, follow_up.Code)
	return
}
