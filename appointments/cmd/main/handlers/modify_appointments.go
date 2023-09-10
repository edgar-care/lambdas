package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func ModifRdv(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id_appointment := chi.URLParam(req, "id")

	if id_appointment == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ID is not defined",
		}, 400)
		return
	}

	// ======================================================= //
	var new_appointment services.RdvInput
	err := json.NewDecoder(req.Body).Decode(&new_appointment)

	lib.CheckError(err)
	appointment, err := services.GetRdvById(new_appointment.Id)

	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to an appointment",
		}, 400)
		return
	}

	if appointment.IdPatient != "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Is already book",
		}, 400)
		return
	}

	rdv, err := services.UpdateRdv(patientID, new_appointment.Id)

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
		RendezVousIDs: append(patient.RendezVousIDs, new_appointment.Id),
	}

	updatedPatient, err := services.UpdatePatient(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	// =============================================================== //

	_, err = services.UpdateRdv("", id_appointment)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	patient, err = services.GetPatientById(patientID)
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

	updatedPatient, err = services.UpdatePatient(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv":     rdv,
		"patient": updatedPatient,
	}, 201)
}
