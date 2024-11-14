package handlers

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
	"net/http"
)

func AddMedicalAntecedent(w http.ResponseWriter, req *http.Request) {

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

	var input edgarlib.CreateNewMedicalAntecedentInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	spew.Dump(input)

	medical_antecedent := edgarlib.AddMedicalAntecedent(input, patientID.ID)
	if medical_antecedent.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medical_antecedent.Err.Error(),
		}, medical_antecedent.Code)
		return
	}

	lib.WriteResponse(w, medical_antecedent.MedicalAntecedents, medical_antecedent.Code)
}
