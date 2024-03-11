package handlers

import (
	"encoding/json"
	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/dashboard"
	"net/http"
)

func CreatePatient(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	var input edgarlib.CreatePatientInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	patient := edgarlib.CreatePatientFormDoctor(input, doctorID)
	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patient":      patient.Patient,
		"medical_info": patient.MedicalInfo,
	}, 201)
}
