package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"net/http"
)

type SenderEmailInput struct {
	Email string `json:"email"`
}

func SenderEmail2FA(w http.ResponseWriter, req *http.Request) {

	var input SenderEmailInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	if !checkerEmail(input.Email) {
		lib.WriteError(w, http.StatusBadRequest, "Invalid email")
		return
	}

	sendingEmail := edgarlib.Email2faAuth(input.Email)
	if sendingEmail.Err != nil {
		lib.WriteError(w, sendingEmail.Code, sendingEmail.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"email": "send",
	}, sendingEmail.Code)
}

func checkerEmail(email string) bool {
	_, patientErr := graphql.GetPatientByEmail(email)
	if patientErr != nil {
		_, doctorErr := graphql.GetDoctorByEmail(email)
		if doctorErr != nil {
			return false
		}
		return true
	}
	return true
}
