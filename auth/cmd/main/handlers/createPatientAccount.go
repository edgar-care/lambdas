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

type CreatePatientInput struct {
	Email string `json:"email"`
}

func CreatePatientAccount(w http.ResponseWriter, req *http.Request) {
	var input CreatePatientInput
	var register services.PatientInput
	var email edgarEmail.Email

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	password := services.GeneratePassword(10)
	register.Email = input.Email
	register.Password = password

	patient, err := services.RegisterPatient(register)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": err.Error(),
		}, 400)
		return
	}
	patient_uuid := uuid.New()
	expire := 43200
	_, err = redis.SetKey(patient_uuid.String(), input.Email, &expire)
	lib.CheckError(err)

	email.To = input.Email
	email.Subject = "Création de votre compte - edgar-sante.fr"
	email.Body = fmt.Sprintf("Votre compte à bien été créé, cliquez ici pour mettre à jour votre mot de passe (app.edgar-sante.fr/reset-password?uuid=%s)", patient_uuid.String())

	err = edgarEmail.SendEmail(email)
	lib.CheckError(err)

	lib.WriteResponse(w, map[string]string{
		"id": patient.Id,
	}, 200)
}
