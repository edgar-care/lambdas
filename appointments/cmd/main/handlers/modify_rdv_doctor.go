package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func UpdateDoctorAppointment(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	appointmentID := chi.URLParam(req, "id")

	if appointmentID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ID is not defined",
		}, 400)
		return
	}

	var new_appointment services.RdvInput
	err := json.NewDecoder(req.Body).Decode(&new_appointment)
	lib.CheckError(err)

	appointment, err := services.GetRdvById(new_appointment.Id)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to an appointment",
		}, 400)
		return
	}

	if appointment.IdPatient != "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Is already book",
		}, 400)
		return
	}

	get_id, err := services.GetRdvById(appointmentID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to an appointment",
		}, 400)
		return
	}

	rdv, err := services.UpdateRdv(get_id.IdPatient, new_appointment.Id)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	var updatePatient services.PatientInput

	patient, err := services.GetPatientById(get_id.IdPatient)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to a patient",
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id:            get_id.IdPatient,
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

	_, err = services.UpdateRdv("", appointmentID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	patient, err = services.GetPatientById(get_id.IdPatient)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to a patient",
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id:            get_id.IdPatient,
		RendezVousIDs: removeElement(patient.RendezVousIDs, appointmentID),
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
