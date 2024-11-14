package handlers

import (
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DeleteMedicalAntecedent(w http.ResponseWriter, req *http.Request) {

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

	deleted_medical_antecedent := edgarlib.DeleteMedicalAntecedent(t, patientID.ID)
	if deleted_medical_antecedent.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": deleted_medical_antecedent.Err.Error(),
		}, deleted_medical_antecedent.Code)
		return
	}

	response := map[string]interface{}{
		"delete": deleted_medical_antecedent.Deleted,
	}

	lib.WriteResponse(w, response, deleted_medical_antecedent.Code)

}
