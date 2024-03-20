package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
)

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
	lib.WriteResponse(w, map[string]interface{}{
		"review": reviewWait.Rdv,
	}, reviewWait.Code)

}
