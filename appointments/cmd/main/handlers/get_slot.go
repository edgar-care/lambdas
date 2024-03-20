package handlers

import (
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/slot"
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

	id := chi.URLParam(req, "id")

	slot := edgarlib.GetSlotById(id, doctorID)

	if slot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": slot.Err.Error(),
		}, slot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot.Slot,
	}, 200)
}

func GetSlots(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	slot := edgarlib.GetSlots(doctorID)

	if slot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": slot.Err.Error(),
		}, slot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot.Slots,
	}, 200)
}
