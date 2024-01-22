package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	"github.com/edgar-care/dashboard/cmd/main/services"
)

func CreatePatient(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	var input services.MedicalInfo

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	url := os.Getenv("URL")
	payload, err := json.Marshal(input.Patient)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert the response body to a string

	var responseStruct struct {
		ID string `json:"id"`
	}

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(responseBody, &responseStruct)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Get the ID value from the struct
	patient := responseStruct.ID
	fmt.Print(patient)

	// =============================================== //

	// Update doctor

	var updatedDoctor services.DoctorInput
	doctor, err := services.GetDoctorById(doctorID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a doctor",
		}, 400)
		return
	}

	updatedDoctor = services.DoctorInput{
		Id:         doctorID,
		PatientIds: append(doctor.PatientIds, patient),
	}
	updatDoctor, err := services.UpdateDoctor(updatedDoctor)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}
	//========================================================================//
	// Add information for the patient in onboarding part
	var updatePatient services.PatientHealthInput
	var updateHealtPatient services.PatientHealthInput

	info, err := services.CreateInfo(input.Info)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	// Add health information for the patient in onboarding part

	lib.CheckError(err)
	health, err := services.CreateHealth(input.Health)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info Health (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	// Update patient
	updatePatient = services.PatientHealthInput{
		Id:                 patient,
		OnboardingInfoID:   info.Id,
		OnboardingHealthID: health.Id,
	}
	updateHealtPatient = services.PatientHealthInput{
		Id:                 patient,
		OnboardingHealthID: health.Id,
	}
	_, err = services.AddOnboardingInfoID(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}
	_, err = services.AddOnboardingHealthID(updateHealtPatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	// Return value

	lib.WriteResponse(w, map[string]interface{}{
		"patient":        patient,
		"patient_health": health,
		"patient_info":   info,
		"update doctor":  updatDoctor,
	}, 200)
}
