package handlers

import (
	"net/http"

	"github.com/edgar-care/slot/cmd/main/lib"
	"github.com/edgar-care/slot/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func DeleteSlot(w http.ResponseWriter, req *http.Request) {
	// Extract doctorID from the request using your authentication middleware
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	slotID := chi.URLParam(req, "id")

	deletedSlot, err := services.DeleteSlotId(slotID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Error deleting slot: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	var updatedDoctor services.DoctorInput
	doctor, err := services.GetDoctorById(doctorID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a doctor",
		}, 400)
		return
	}

	updatedDoctor = services.DoctorInput{
		Id:      doctorID,
		SlotIDs: removeElement(doctor.SlotIDs, slotID),
	}
	updatDoctor, err := services.UpdateDoctor(updatedDoctor)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	// Respond with a message indicating whether the slot was deleted successfully
	lib.WriteResponse(w, map[string]interface{}{
		"message": "Slot deleted successfully",
		"deleted": deletedSlot, // Assuming 'deleted' is a suitable name for the boolean value
		"update":  updatDoctor,
	}, http.StatusOK)
}

func removeElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
