package handlers

import (
	"encoding/json"
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func EditMedicalAntecedent(w http.ResponseWriter, req *http.Request) {

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

	var input edgarlib.UpdateMedicalFolderPatientInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	edit_medical_antecedent := edgarlib.UpdateMedicalAntecedent(patientID.ID, input, t)
	if edit_medical_antecedent.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": edit_medical_antecedent.Err.Error(),
		}, edit_medical_antecedent.Code)
		return
	}

	lib.WriteResponse(w, edit_medical_antecedent.UpdatedAntecedent, edit_medical_antecedent.Code)

}
