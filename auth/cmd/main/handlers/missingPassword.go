package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	edgarEmail "github.com/edgar-care/edgarlib/email"
	"github.com/edgar-care/edgarlib/redis"
	"github.com/google/uuid"

	"net/http"
)

type MissingPasswordInput struct {
	Email string `json:"email"`
}

func MissingPassword(w http.ResponseWriter, req *http.Request) {
	var input MissingPasswordInput
	var email edgarEmail.Email

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	_, err = services.GetPatientByEmail(input.Email)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "No patient corresponds to this email",
		}, 400)
		return
	}
	patient_uuid := uuid.New()
	expire := 600
	_, err = redis.SetKey(patient_uuid.String(), input.Email, &expire)
	lib.CheckError(err)

	email.To = input.Email
	email.Subject = "Réinitialisation de votre mot de passe"
	email.Body = fmt.Sprintf("Pour réinitialiser votre mot de passe, cliquez ici (app.edgar-sante.fr/reset-password?uuid=%s)", patient_uuid.String())

	err = edgarEmail.SendEmail(email)
	lib.CheckError(err)

	lib.WriteResponse(w, map[string]string{}, 200)
}
