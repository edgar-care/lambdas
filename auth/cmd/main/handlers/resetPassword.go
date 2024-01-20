package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	"github.com/edgar-care/edgarlib/redis"
	"github.com/jinzhu/copier"
	"net/http"
	"strings"
)

type ResetPasswordInput struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func ResetPassword(w http.ResponseWriter, req *http.Request) {
	var input ResetPasswordInput
	var updatePatient services.PatientUpdateInput
	uuid := req.URL.Query().Get("uuid")
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	if uuid == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "uuid has to be provided",
		}, 403)
		return
	}
	value, err := redis.GetKey(uuid)
	value = strings.Replace(value, "\n", "", -1)
	if value == "" || err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "uuid is expired",
		}, 403)
		return
	}

	patient, err := services.GetPatientByEmail(input.Email)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "No patient corresponds to this email",
		}, 400)
		return
	}

	patient.Password = lib.HashPassword(input.NewPassword)
	fmt.Printf("patient:")
	spew.Dump(patient)

	err = copier.Copy(&updatePatient, &patient)
	lib.CheckError(err)

	services.UpdatePatient(updatePatient)

	lib.WriteResponse(w, map[string]string{}, 200)
}
