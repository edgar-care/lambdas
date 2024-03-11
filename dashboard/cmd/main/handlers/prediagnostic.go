package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
)

type RdvSessionCombined struct {
	ID                string                  `json:"id"`
	DoctorID          string                  `json:"doctor_id"`
	PatientID         string                  `json:"id_patient"`
	StartDate         int                     `json:"start_date"`
	EndDate           int                     `json:"end_date"`
	CancelationReason *string                 `json:"cancelation_reason"`
<<<<<<< HEAD
=======
	HealthMethod      *string                 `json:"health_method"`
>>>>>>> 2a04d28 (fix(medicalFolder): wrong lambda use)
	AppointmentStatus model.AppointmentStatus `json:"appointment_status"`
	SessionID         string                  `json:"session_id"`
	Diseases          []model.SessionDiseases `json:"diseases"`
	Fiability         float64                 `json:"fiability"`
	Symptoms          []model.SessionSymptom  `json:"symptoms"`
	Logs              []graphql.LogsInput     `json:"logs"`
	Alerts            []model.Alert           `json:"alerts"`
}

func RevPreDiagnostic(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id_rdv := chi.URLParam(req, "id")
	var input edgarlib.ReviewInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	review := edgarlib.ValidateRdv(id_rdv, input)
	if review.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": review.Err.Error(),
		}, review.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"review": review.Rdv,
	}, review.Code)
}

func GetPreDignosticWait(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	reviewWait := edgarlib.GetWaitingReview(doctorID)
	if reviewWait.Err != nil {
		lib.WriteResponse(w, map[string]interface{}{
			"message": reviewWait.Err.Error(),
		}, reviewWait.Code)
		return
	}
	var responseList []RdvSessionCombined
	for _, rdvSession := range reviewWait.RdvWithSession {
		responseList = append(responseList, RdvSessionCombined{
			ID:                rdvSession.Rdv.ID,
			DoctorID:          rdvSession.Rdv.DoctorID,
			PatientID:         rdvSession.Rdv.IDPatient,
			StartDate:         rdvSession.Rdv.StartDate,
			EndDate:           rdvSession.Rdv.EndDate,
			CancelationReason: rdvSession.Rdv.CancelationReason,
<<<<<<< HEAD
=======
			HealthMethod:      rdvSession.Rdv.HealthMethod,
>>>>>>> 2a04d28 (fix(medicalFolder): wrong lambda use)
			AppointmentStatus: rdvSession.Rdv.AppointmentStatus,
			SessionID:         rdvSession.Rdv.SessionID,
			Diseases:          rdvSession.Session.Diseases,
			Fiability:         rdvSession.Session.Fiability,
			Symptoms:          rdvSession.Session.Symptoms,
			Logs:              rdvSession.Session.Logs,
			Alerts:            rdvSession.Session.Alerts,
		})
	}

	lib.WriteResponse(w, responseList, reviewWait.Code)

}
