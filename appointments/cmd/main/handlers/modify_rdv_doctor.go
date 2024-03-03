package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
	"github.com/go-chi/chi/v5"
)

type UpdateRdvDoctorInput struct {
	ID string `json:"id"`
}

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

	var new_appointment UpdateRdvDoctorInput
	err := json.NewDecoder(req.Body).Decode(&new_appointment)
	lib.CheckError(err)

	appointment := edgarlib.UpdateDoctorAppointment(new_appointment.ID, appointmentID)
	if appointment.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": appointment.Err.Error(),
		}, appointment.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": appointment.Rdv,
	}, 201)
}
