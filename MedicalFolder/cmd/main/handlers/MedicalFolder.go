package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/medical_folder"
)

func AddMedicalInfo(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateMedicalInfoInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medical := edgarlib.CreateMedicalInfo(input, patientID)
	if medical.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medical.Err.Error(),
		}, medical.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"MedicalFolder": medical.MedicalInfo,
	}, 201)
}
