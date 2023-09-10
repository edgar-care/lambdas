package handlers

import (
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func BookRdv(w http.ResponseWriter, req *http.Request) {

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

	appointment, err := services.GetRdvById(id_appointment)

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

	rdv, err := services.UpdateRdv(patientID, id_appointment)

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
		RendezVousIDs: append(patient.RendezVousIDs, id_appointment),
	}

	updatedPatient, err := services.UpdatePatient(updatePatient)
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

func GetRdvPatient(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	rdv, err := services.GetRdvById(t)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	if rdv.IdPatient != patientID {
		lib.WriteResponse(w, map[string]string{
			"message": "You can't access to this appointment",
		}, 403)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv,
	}, 201)
}

func GetRdv(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	rdv, err := services.GetAllRdv(patientID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv,
	}, 201)
}
