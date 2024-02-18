package handlers

import (
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
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

	appointment := edgarlib.BookAppointment(id_appointment, patientID)

	if appointment.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": appointment.Err.Error(),
		}, appointment.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": appointment,
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

	rdv := edgarlib.GetRdvPatient(t, patientID)

	if rdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": rdv.Err.Error(),
		}, rdv.Code)
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
	rdv := edgarlib.GetRdv(patientID)

	if rdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": rdv.Err.Error(),
		}, rdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv.Rdv,
	}, 201)
}
