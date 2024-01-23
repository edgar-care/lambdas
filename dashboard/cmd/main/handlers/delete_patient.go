package handlers

import (
	"fmt"
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	"github.com/edgar-care/dashboard/cmd/main/services"
	edgarEmail "github.com/edgar-care/edgarlib/email"
	"github.com/edgar-care/edgarlib/redis"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func DeletePatientHandler(w http.ResponseWriter, req *http.Request) {
	var email edgarEmail.Email

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	PatientID := chi.URLParam(req, "id")

	// stock redis
	patient_uuid := uuid.New()
	expire := 2592000
	_, err := redis.SetKey(patient_uuid.String(), PatientID, &expire)
	lib.CheckError(err)

	go watchExpiration(w, patient_uuid.String(), PatientID, doctorID)

	patient, err := services.GetPatientById(PatientID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a patient",
		}, 400)
		return
	}

	email.To = patient.Email
	email.Subject = "Votre compte - edgar-sante.fr va étre supprimer dans 30 jour"
	email.Body = fmt.Sprintf("Si cela n'est pas normale veuillez contacter votre medecin", patient_uuid.String())

	err = edgarEmail.SendEmail(email)
	lib.CheckError(err)

	lib.WriteResponse(w, map[string]interface{}{
		"deleted": email.Subject, // Assuming 'deleted' is a suitable name for the boolean value
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

func watchExpiration(w http.ResponseWriter, patientUUID string, PatientID string, doctorID string) {
	// Vérification périodique de l'expiration
	//time.Sleep(24 * time.Hour) // Changer cela en la fréquence souhaitée

	// Récupération de l'ID du patient depuis Redis
	patientID, err := redis.GetKey(patientUUID)
	if err != nil {
		// Gérer l'erreur
		fmt.Println("Error retrieving patient ID from Redis:", err)
		return
	}

	// Si l'ID du patient n'est pas vide, cela signifie que l'expiration n'a pas encore eu lieu
	if patientID != "" {
		// Envoyer un e-mail périodique
		sendPeriodicEmail(patientID, patientUUID)
	} else {
		// L'expiration a eu lieu, arrêter la surveillance
		deletepatient, err := services.DeleteSlotId(PatientID)
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
			Id:         doctorID,
			PatientIds: removeElement(doctor.PatientIds, PatientID),
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
			"message": "Patient deleted successfully",
			"deleted": deletepatient, // Assuming 'deleted' is a suitable name for the boolean value
			"update":  updatDoctor,
		}, http.StatusOK)
		return
	}
}

func sendPeriodicEmail(patientID, patientUUID string) {
	patient, err := services.GetPatientById(patientID)
	if err != nil {
		// Gérer l'erreur
		fmt.Println("Error getting patient information:", err)
		return
	}

	var email edgarEmail.Email
	email.To = patient.Email
	email.Subject = "Rappel - Suppression du compte dans 30 jours"
	email.Body = fmt.Sprintf("Si cela n'est pas normal, veuillez contacter votre médecin. UUID: %s", patientUUID)

	err = edgarEmail.SendEmail(email)
	if err != nil {
		// Gérer l'erreur
		fmt.Println("Error sending email:", err)
	}
}
