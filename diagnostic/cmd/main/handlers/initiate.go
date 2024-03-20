package handlers

import (
	edgarlib "github.com/edgar-care/edgarlib/auth"
	edgar_diag "github.com/edgar-care/edgarlib/diagnostic"

	edgarhttp "github.com/edgar-care/edgarlib/http"
	"net/http"
)

func Initiate(w http.ResponseWriter, req *http.Request) {
	patientID := edgarlib.AuthMiddleware(w, req)
	if patientID == "" {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	resp := edgar_diag.Initiate(patientID)

	if resp.Err != nil {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"sessionId": resp.Id,
	}, resp.Code)
}
