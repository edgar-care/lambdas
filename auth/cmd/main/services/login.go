package services

import (
	"github.com/edgar-care/auth/cmd/main/lib"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(input LoginInput, t string) (string, error) {
	var doctor interface{}
	var patient interface{}
	var token string
	var err error
	if t == "d" {
		doctor, err = GetDoctorByEmail(input.Email)
	} else {
		patient, err = GetPatientByEmail(input.Email)
	}
	if err != nil {
		return "Could not find user: " + err.Error(), err
	}

	if !(t == "d" && lib.CheckPassword(input.Password, doctor.(Doctor).Password)) &&
		!(t == "p" && lib.CheckPassword(input.Password, patient.(Patient).Password)) {
		return "Username and password mismatch.", err
	}

	if t == "d" {
		token, err = lib.CreateToken(map[string]interface{}{
			"doctor": doctor,
		})
	} else {
		token, err = lib.CreateToken(map[string]interface{}{
			"patient": patient,
		})
	}
	return token, err
}
