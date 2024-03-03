package services

import (
	"fmt"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/auth"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

func RegisterDoctor(input edgarlib.DoctorInput) (model.Doctor, error) {
	input.Password = lib.HashPassword(input.Password)
	doctor, err := edgarlib.RegisterDoctor(input.Email, input.Password, input.Name, input.Firstname, input.Address)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	return doctor, nil
}

func RegisterAndLoginDoctor(input edgarlib.DoctorInput) (string, int, error) {
	doctor, err := edgarlib.RegisterAndLoginDoctor(input)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"doctor": doctor,
	})
	return token, 200, err
}

func RegisterPatient(input edgarlib.PatientInput) (model.Patient, error) {
	input.Password = lib.HashPassword(input.Password)
	patient, err := RegisterPatient(input.Email, input.Password)
	if err != nil {
		return patient, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return patient, nil
}

func RegisterAndLoginPatient(input edgarlib.PatientInput) (string, int, error) {
	patient, err := edgarlib.RegisterAndLoginPatient(input)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"patient": patient,
	})
	return token, 200, err
}

func RegisterAdmin(input edgarlib.AdminInput) (model.Admin, error) {
	if lib.VerifyToken(input.Token) == false {
		return model.Admin{}, fmt.Errorf("Unable to create account: Invalid Token")
	}
	input.Password = lib.HashPassword(input.Password)
	admin, err := edgarlib.RegisterAdmin(input)
	if err != nil {
		return admin, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return admin, nil
}

func RegisterAndLoginAdmin(input edgarlib.AdminInput) (string, int, error) {
	admin, err := edgarlib.RegisterAndLoginAdmin(input)
	if err != nil {
		return "", 400, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"admin": admin,
	})
	return token, 200, err
}
