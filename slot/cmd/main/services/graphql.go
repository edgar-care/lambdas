package services

import (
	"context"
	"fmt"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Slot struct {
	Id             string `json:"id"`
	DoctorID       string `json:"doctor_id"`
	AppointmentsID string `json:"appointment_id"`
	StartDate      int    `json:"start_date"`
	EndDate        int    `json:"end_date"`
}

type SlotOutput struct {
	Id             string  `json:"id"`
	DoctorID       *string `json:"doctor_id"`
	AppointmentsID *string `json:"appointment_id"`
	StartDate      *int    `json:"start_date"`
	EndDate        *int    `json:"end_date"`
}

type SlotInput struct {
	Id             string `json:"id"`
	DoctorID       string `json:"doctor_id"`
	AppointmentsID string `json:"appointment_id"`
	StartDate      int    `json:"start_date"`
	EndDate        int    `json:"end_date"`
}

type Doctor struct {
	Id      string   `json:"id"`
	SlotIDs []string `json:"slot_ids"`
}

type DoctorInput struct {
	Id      string   `json:"id"`
	SlotIDs []string `json:"slot_ids"`
}

type DoctorOutput struct {
	Id      *string   `json:"id"`
	SlotIDs *[]string `json:"slot_ids"`
}

/**************** GraphQL types *****************/

type createSlotResponse struct {
	Content SlotOutput `json:"createSlot"`
}

type getOneSlotByIdResponse struct {
	Content SlotOutput `json:"getSlotById"`
}

type getAllSlotResponse struct {
	Content []SlotOutput `json:"getDoctorSlot"`
}

type deleteSlotResponse struct {
	Content bool `json:"deleteSlot"`
}

type updateDoctorResponse struct {
	Content DoctorOutput `json:"updateDoctor"`
}

/*************** Implementations *****************/

func CreateSlots(slotin SlotInput, id string) (Slot, error) {
	var slot createSlotResponse
	var resp Slot

	query := `mutation createSlot($doctor_id: String!, $start_date: Int!, $end_date: Int!) {
		createSlot(doctor_id:$doctor_id, start_date:$start_date, end_date:$end_date) {
                    id,
					doctor_id,
					start_date,
					end_date,
					appointment_id
                }
            }`
	err := Query(query, map[string]interface{}{
		"doctor_id":      id,
		"start_date":     slotin.StartDate,
		"end_date":       slotin.EndDate,
		"appointment_id": slotin.AppointmentsID,
	}, &slot)
	_ = copier.Copy(&resp, &slot.Content)
	fmt.Print(resp)
	return resp, err
}

func GetRdvById(id string) (Slot, error) {
	var oneslot getOneSlotByIdResponse
	var resp Slot
	query := `query getRdvById($id: String!) {
                getRdvById(id: $id) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &oneslot)
	_ = copier.Copy(&resp, &oneslot.Content)
	return resp, err
}

func GetAllSlot(id string) ([]Slot, error) {
	var allrdv getAllSlotResponse
	var resp []Slot
	query := `query getDoctorSlot($doctor_id: String!){
                getDoctorSlot(doctor_id: $doctor_id) {
                    id,
					doctor_id,
					start_date,
					end_date
                }
            }`
	err := Query(query, map[string]interface{}{
		"doctor_id": id,
	}, &allrdv)
	_ = copier.Copy(&resp, &allrdv.Content)
	return resp, err
}

func DeleteSlotId(id string) (Slot, error) {
	var slotdelete deleteSlotResponse
	var resp Slot
	query := `mutation deleteSlot($id: String!) {
		deleteSlot(id:$id)
	}`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &slotdelete)
	_ = copier.Copy(&resp, &slotdelete.Content)
	return resp, err
}

// ======================================================================================== //

func GetSlotDoctorById(id string) (Slot, error) {
	var oneslot getOneSlotByIdResponse
	var resp Slot
	query := `query getSlotById($id: String!){
		getSlotById(id:$id) {
                    id,
					doctor_id,
					start_date,
					end_date
                }
            }`
	err := Query(query, map[string]interface{}{
		"id": id,
	}, &oneslot)
	_ = copier.Copy(&resp, &oneslot.Content)
	return resp, err
}

// ============================================================================================== //
// Doctor
func UpdateDoctor(updateDoctor DoctorInput) (Doctor, error) {
	var doctor updateDoctorResponse
	var resp Doctor
	query := `mutation updateDoctor($id: String!, $slot_ids: [String]) {
		updateDoctor(id:$id, slot_ids:$slot_ids) {
                    id,
					slot_ids
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":       updateDoctor.Id,
		"slot_ids": updateDoctor.SlotIDs,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func GetDoctorById(id string) (Doctor, error) {
	var doctor updateDoctorResponse
	var resp Doctor
	query := `query getDoctorById($id: String!) {
                getDoctorById(id: $id) {
                    id,
					slot_ids
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func createClient() *graphql.Client {
	return graphql.NewClient(os.Getenv("GRAPHQL_URL"))
}

func Query(query string, variables map[string]interface{}, respData interface{}) error {
	var request = graphql.NewRequest(query)
	var ctx = context.Background()
	for key, value := range variables {
		request.Var(key, value)
	}
	request.Header.Set(os.Getenv("API_KEY"), os.Getenv("API_KEY_VALUE"))
	err := createClient().Run(ctx, request, respData)
	return err
}
