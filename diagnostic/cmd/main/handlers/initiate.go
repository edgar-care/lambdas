package handlers

import (
	"fmt"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	edgarauth "github.com/edgar-care/edgarlib/v2/auth"
	edgar_diag "github.com/edgar-care/edgarlib/v2/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/v2/http"
	"net/http"
)

func Initiate(w http.ResponseWriter, req *http.Request) {
	patientID := edgarauth.AuthMiddlewarePatient(w, req)
	if patientID.Code == 409 || patientID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": patientID.Err.Error(),
		}, patientID.Code)
		return
	}
	if patientID.ID == "" {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	fmt.Print(patientID)

	resp := edgar_diag.Initiate(patientID.ID)

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
