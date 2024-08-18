package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/follow_treatment"
	"github.com/edgar-care/treatment_follow_up/cmd/main/lib"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DeleteFollowTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := authlib.AuthMiddlewarePatient(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	follow_delete := edgarlib.Delete_follow_up(t, patientID)

	if follow_delete.Err != nil {
		lib.WriteError(w, follow_delete.Code, follow_delete.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"delete": follow_delete.Deleted,
	}, follow_delete.Code)
}
