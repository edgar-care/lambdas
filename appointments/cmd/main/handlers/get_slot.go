package handlers

import (
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func GetSlotId(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	slot, err := services.GetRdvById(t)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}
	if slot.DoctorID != doctorID {
		lib.WriteResponse(w, map[string]string{
			"message": "You can't access to this appointment",
		}, 403)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot,
	}, 201)
}

func GetSlots(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	slot, err := services.GetAllRdvDoctor(doctorID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot,
	}, 201)
}
