package services

import (
	"errors"

	"github.com/edgar-care/auth/cmd/main/lib"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(input LoginInput, t string) (string, error) {
	var doctor interface{}
	var admin interface{}
	var patient interface{}
	var token string
	var err error
	if t == "d" {
		doctor, err = GetDoctorByEmail(input.Email)
	} else if t == "a" {
		admin, err = GetAdminByEmail(input.Email)
	} else {
		patient, err = GetPatientByEmail(input.Email)
	}
	if err != nil {
		return "Could not find user: " + err.Error(), err
	}

	if !(t == "d" && lib.CheckPassword(input.Password, doctor.(Doctor).Password)) &&
		!(t == "a" && lib.CheckPassword(input.Password, admin.(Admin).Password)) &&
		!(t == "p" && lib.CheckPassword(input.Password, patient.(Patient).Password)) {
		return "Username and password mismatch.", errors.New("username and password mismatch")
	}

	if t == "d" {
		token, err = lib.CreateToken(map[string]interface{}{
			"doctor": doctor,
		})
	} else if t == "a" {
		token, err = lib.CreateToken(map[string]interface{}{
			"admin": admin,
		})
	} else {
		token, err = lib.CreateToken(map[string]interface{}{
			"patient": patient,
		})
	}
	return token, err
}
