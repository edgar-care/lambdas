package handlers

import (
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetMedicalAntecedentByID(w http.ResponseWriter, req *http.Request) {

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

	get_medical_antecedent := edgarlib.GetMedicalAntecedentById(t, patientID.ID)
	if get_medical_antecedent.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": get_medical_antecedent.Err.Error(),
		}, get_medical_antecedent.Code)
		return
	}

	lib.WriteResponse(w, get_medical_antecedent.MedicalAntecedent, get_medical_antecedent.Code)
}

func GetMedicalAntecedents(w http.ResponseWriter, req *http.Request) {

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

	medical_antecedents := edgarlib.GetMedicalAntecedents(patientID.ID)
	if medical_antecedents.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medical_antecedents.Err.Error(),
		}, medical_antecedents.Code)
		return
	}

	lib.WriteResponse(w, medical_antecedents.MedicalAntecedents, medical_antecedents.Code)
}
