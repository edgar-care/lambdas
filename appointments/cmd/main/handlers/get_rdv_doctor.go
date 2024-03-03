package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
)

func GetDoctorAppointment(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	appointmentID := chi.URLParam(req, "id")

	rdv := edgarlib.GetDoctorAppointment(appointmentID, doctorID)

	if rdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": rdv.Err.Error(),
		}, rdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv.Appointment,
	}, 200)
}

func GetAllDoctorAppointments(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	appointments := edgarlib.GetAllDoctorAppointment(doctorID)

	if appointments.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": appointments.Err.Error(),
		}, appointments.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"appointments": appointments.Slots,
	}, 200)
}
