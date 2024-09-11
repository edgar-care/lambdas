package handlers

import (
	"net/http"
	"strconv"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
)

type DoctorResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Address   struct {
		Street  string `json:"street"`
		ZipCode string `json:"zip_code"`
		Country string `json:"country"`
		City    string `json:"city"`
	} `json:"address"`
}

func GetDoctorId(w http.ResponseWriter, req *http.Request) {

	t := chi.URLParam(req, "id")

	doctor := edgarlib.GetDoctorById(t)
	if doctor.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": doctor.Err.Error(),
		}, doctor.Code)
		return
	}

	doctorResponse := DoctorResponse{
		ID:        doctor.Doctor.ID,
		Email:     doctor.Doctor.Email,
		Name:      doctor.Doctor.Name,
		Firstname: doctor.Doctor.Firstname,
		Address: struct {
			Street  string `json:"street"`
			ZipCode string `json:"zip_code"`
			Country string `json:"country"`
			City    string `json:"city"`
		}{
			Street:  doctor.Doctor.Address.Street,
			ZipCode: doctor.Doctor.Address.ZipCode,
			Country: doctor.Doctor.Address.Country,
			City:    doctor.Doctor.Address.City,
		},
	}

	lib.WriteResponse(w, map[string]interface{}{
		"Doctors": doctorResponse,
	}, 200)
}

func GetDoctors(w http.ResponseWriter, req *http.Request) {
	var doctors edgarlib.GetDoctorsResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		doctors = edgarlib.GetDoctors(0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			doctors = edgarlib.GetDoctors(0, 0)
		} else {
			doctors = edgarlib.GetDoctors(number_page, number_size)
		}
	}

	if doctors.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": doctors.Err.Error(),
		}, doctors.Code)
		return
	}
	var doctorsResponse []DoctorResponse
	for _, doctor := range doctors.Doctors {
		doctorResponse := DoctorResponse{
			ID:        doctor.ID,
			Email:     doctor.Email,
			Name:      doctor.Name,
			Firstname: doctor.Firstname,
			Address: struct {
				Street  string `json:"street"`
				ZipCode string `json:"zip_code"`
				Country string `json:"country"`
				City    string `json:"city"`
			}{
				Street:  doctor.Address.Street,
				ZipCode: doctor.Address.ZipCode,
				Country: doctor.Address.Country,
				City:    doctor.Address.City,
			},
		}
		doctorsResponse = append(doctorsResponse, doctorResponse)
	}

	lib.WriteResponse(w, map[string]interface{}{
		"Doctors": doctorsResponse,
	}, 200)
}
