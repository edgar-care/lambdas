package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
)

func DeleteRdv(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id_appointment := chi.URLParam(req, "id")
	_, err := services.UpdateRdv("", id_appointment)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	var updatePatient services.PatientInput

	patient, err := services.GetPatientById(patientID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a patient",
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id:            patientID,
		RendezVousIDs: removeElement(patient.RendezVousIDs, id_appointment),
	}

	updatedPatient, err := services.UpdatePatient(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"updated_patient": updatedPatient,
	}, 201)
}

func removeElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
