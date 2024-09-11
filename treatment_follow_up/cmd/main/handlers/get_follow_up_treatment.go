package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"

	edgarlib "github.com/edgar-care/edgarlib/v2/follow_treatment"
	"github.com/edgar-care/treatment_follow_up/cmd/main/lib"
)

func GetFollowTreatment(w http.ResponseWriter, req *http.Request) {

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

	t := chi.URLParam(req, "id")

	follow_get := edgarlib.GetTreatmentFollowUpById(t)
	if follow_get.Err != nil {
		lib.WriteError(w, follow_get.Code, follow_get.Err.Error())
		return
	}

	lib.WriteResponse(w, follow_get.TreatmentFollowUp, follow_get.Code)
}

func GetfFollowsTreatments(w http.ResponseWriter, req *http.Request) {

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
	check_account := authlib.CheckAccountEnable(patientID.ID)
	if check_account.Code == 409 {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authorized, this account is disable",
		}, 409)
		return
	}

	follow_up := edgarlib.GetTreatmentFollowUp(patientID.ID)
	if follow_up.Err != nil {
		lib.WriteError(w, follow_up.Code, follow_up.Err.Error())
		return
	}

	lib.WriteResponse(w, follow_up.TreatmentFollowUps, follow_up.Code)
	return
}
