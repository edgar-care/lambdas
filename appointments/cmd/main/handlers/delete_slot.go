package handlers

import (
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/slot"
	"github.com/go-chi/chi/v5"
)

func DeleteSlot(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	slotID := chi.URLParam(req, "id")

	if slotID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ID is not defined",
		}, 400)
		return
	}

	deletedSlot := edgarlib.DeleteSlot(slotID, doctorID)
	if deletedSlot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": deletedSlot.Err.Error(),
		}, deletedSlot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"message": "Slot deleted successfully",
		"deleted": deletedSlot.Deleted,
	}, http.StatusOK)

}
