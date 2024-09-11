package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/appointment"
)

type RdvSessionCombined struct {
	ID                string  `json:"id"`
	DoctorID          string  `json:"doctor_id"`
	PatientID         string  `json:"id_patient"`
	StartDate         int     `json:"start_date"`
	EndDate           int     `json:"end_date"`
	CancelationReason *string `json:"cancelation_reason"`

	HealthMethod *string `json:"health_method"`

	AppointmentStatus model.AppointmentStatus `json:"appointment_status"`
	SessionID         string                  `json:"session_id"`
	Diseases          []model.SessionDiseases `json:"diseases"`
	Fiability         float64                 `json:"fiability"`
	Symptoms          []model.SessionSymptom  `json:"symptoms"`
	Logs              []model.LogsInput       `json:"logs"`
	Alerts            []model.Alert           `json:"alerts"`
}

type ReturnStruct struct {
	Review []RdvSessionCombined `json:"review"`
}

func RevPreDiagnostic(w http.ResponseWriter, req *http.Request) {
	doctorID := authlib.AuthMiddlewareDoctor(w, req)
	if doctorID.Code == 409 || doctorID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": doctorID.Err.Error(),
		}, doctorID.Code)
		return
	}
	if doctorID.ID == "" {
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
	doctorID := authlib.AuthMiddlewareDoctor(w, req)
	if doctorID.Code == 409 || doctorID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": doctorID.Err.Error(),
		}, doctorID.Code)
		return
	}
	if doctorID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var reviewWait edgarlib.GetWaitingReviewResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		reviewWait = edgarlib.GetWaitingReview(doctorID.ID, 0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			reviewWait = edgarlib.GetWaitingReview(doctorID.ID, 0, 0)
		} else {
			reviewWait = edgarlib.GetWaitingReview(doctorID.ID, number_page, number_size)
		}
	}
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
			HealthMethod:      rdvSession.Rdv.HealthMethod,

			AppointmentStatus: rdvSession.Rdv.AppointmentStatus,
			SessionID:         rdvSession.Rdv.SessionID,
			Diseases:          rdvSession.Session.Diseases,
			Fiability:         rdvSession.Session.Fiability,
			Symptoms:          rdvSession.Session.Symptoms,
			Logs:              rdvSession.Session.Logs,
			Alerts:            rdvSession.Session.Alerts,
		})
	}

	lib.WriteResponse(w, ReturnStruct{responseList}, reviewWait.Code)

}
