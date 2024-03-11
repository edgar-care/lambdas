package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
)

// func AllRdv(w http.ResponseWriter, req *http.Request) {

// 	patientID := lib.AuthMiddleware(w, req)
// 	if patientID == "" {
// 		lib.WriteResponse(w, map[string]string{
// 			"message": "Not authenticated",
// 		}, 401)
// 		return
// 	}

// 	var input services.RdvInput
// 	var updatePatient services.PatientInput
// 	err := json.NewDecoder(req.Body).Decode(&input)

// 	lib.CheckError(err)
// 	rdv, err := services.GetAllRdv()
	
// 	if err != nil {
// 		lib.WriteResponse(w, map[string]string{
// 			"message": "Invalid input: " + err.Error(),
// 		}, 400)
// 		return
// 	}

// 	updatePatient = services.PatientInput{
// 		Id: patientID,
// 		RendezVousID: rdv.Id,
// 	}
// 	patient, err := services.AddAllRdvID(updatePatient)
// 	if err != nil {
// 		lib.WriteResponse(w, map[string]string{
// 			"message": "Update Failed " + err.Error(),
// 		}, 500)
// 		return
// 	}


// 	lib.WriteResponse(w, map[string]interface{}{
// 		"all_rdv": rdv,
// 		"patient": patient,
// 	}, 201)
// }

func OneRdv(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input services.RdvInput
	var updatePatient services.PatientInput
	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	rdv, err := services.GetOneRdvById(input)
	
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id: patientID,
		RendezVousID: rdv.Id,
	}
	patient, err := services.AddAllRdvID(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}


	lib.WriteResponse(w, map[string]interface{}{
		"one_rdv": rdv,
		"patient": patient,
	}, 201)
}