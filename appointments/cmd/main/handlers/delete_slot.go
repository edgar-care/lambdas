package handlers

import (
	"fmt"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
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

	// CONTROL ID PATIENT
	check_id, err := services.GetRdvById(slotID)

	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a slot",
		}, 400)
		return
	}

	if check_id.IdPatient != "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Is already book, you you cannot delete this slot",
		}, 400)
		return
	}
	//DELETE

	deletedSlot, err := services.DeleteSlotId(slotID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Error deleting slot: " + err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	fmt.Print("hello")

	// UPDATE DOCTOR
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
		Id:            doctorID,
		RendezVousIDs: remElement(doctor.RendezVousIDs, slotID),
	}

	updateDoctor, err := services.UpdateDoctor(updatedDoctor)
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
		"update":  updateDoctor,
	}, http.StatusOK)

}

func remElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
