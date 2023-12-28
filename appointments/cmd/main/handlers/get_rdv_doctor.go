package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
)

// Méthode GET pour récupérer un rendez-vous du docteur
func GetDoctorAppointment(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	appointmentID := chi.URLParam(req, "id")

	rdv, err := services.GetRdvById(appointmentID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	// Vérifiez si le rendez-vous appartient au docteur
	if rdv.DoctorID != doctorID {
		lib.WriteResponse(w, map[string]string{
			"message": "You can't access to this appointment",
		}, 403)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv,
	}, 200)
}

// Méthode GET pour récupérer tous les rendez-vous du docteur
func GetAllDoctorAppointments(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	appointments, err := services.GetAllRdvDoctor(doctorID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"appointments": appointments,
	}, 200)
}
