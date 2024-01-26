package services

import (
	"fmt"
	"github.com/edgar-care/auth/cmd/main/lib"
	"net/http"
)

func RegisterDoctor(input DoctorInput) (Doctor, error) {
	input.Password = lib.HashPassword(input.Password)
	doctor, err := CreateDoctor(input)
	if err != nil {
		return doctor, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return doctor, nil
}

func RegisterAndLoginDoctor(input DoctorInput) (string, int, error) {
	doctor, err := RegisterDoctor(input)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"doctor": doctor,
	})
	return token, 200, err
}

func RegisterPatient(input PatientInput) (Patient, error) {
	input.Password = lib.HashPassword(input.Password)
	patient, err := CreatePatient(input)
	if err != nil {
		return patient, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return patient, nil
}

func RegisterAndLoginPatient(input PatientInput) (string, int, error) {
	patient, err := RegisterPatient(input)
	if err != nil {
		return "", http.StatusBadRequest, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"patient": patient,
	})
	return token, 200, err
}

func RegisterAdmin(input AdminInput) (Admin, error) {
	if lib.VerifyToken(input.Token) == false {
		return Admin{}, fmt.Errorf("Unable to create account: Invalid Token")
	}
	input.Password = lib.HashPassword(input.Password)
	admin, err := CreateAdmin(input)
	if err != nil {
		return admin, fmt.Errorf("Unable to create account: %s", err.Error())
	}
	return admin, nil
}

func RegisterAndLoginAdmin(input AdminInput) (string, int, error) {
	admin, err := RegisterAdmin(input)
	if err != nil {
		return "", 400, err
	}
	token, err := lib.CreateToken(map[string]interface{}{
		"admin": admin,
	})
	return token, 200, err
}
