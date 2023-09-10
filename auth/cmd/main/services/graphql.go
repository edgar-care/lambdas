package services

import (
	"context"
	"os"

	"github.com/jinzhu/copier"
	"github.com/machinebox/graphql"
)

/********** Types ***********/

type Patient struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	Sex      string `json:"sex"`
}

type PatientOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	LastName *string `json:"lastName"`
	Email    *string `json:"email"`
	Age      *int    `json:"age"`
	Height   *int    `json:"height"`
	Weight   *int    `json:"weight"`
	Sex      *string `json:"sex"`
}

type PatientInput struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	Sex      string `json:"sex"`
}

type Doctor struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Admin struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type DoctorOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	LastName *string `json:"lastName"`
	Email    *string `json:"email"`
	Address  *string `json:"address"`
}

type AdminOutput struct {
	Id       *string `json:"id"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	LastName *string `json:"lastName"`
	Email    *string `json:"email"`
}

type DoctorInput struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type AdminInput struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

/**************** GraphQL types *****************/

type getPatientByEmailResponse struct {
	Content PatientOutput `json:"getPatientByEmail"`
}

type getPatientByIdResponse struct {
	Content PatientOutput `json:"getPatientById"`
}

type createPatientResponse struct {
	Content PatientOutput `json:"createPatient"`
}

type getDoctorByIdResponse struct {
	Content DoctorOutput `json:"getDoctorById"`
}

type getDoctorByEmailResponse struct {
	Content DoctorOutput `json:"getDoctorByEmail"`
}

type createDoctorResponse struct {
	Content DoctorOutput `json:"createDoctor"`
}

type getAdminByEmailResponse struct {
	Content AdminOutput `json:"getAdminByEmail"`
}

type createAdminResponse struct {
	Content AdminOutput `json:"createAdmin"`
}

/*************** Implementations *****************/

func GetPatientById(id string) (Patient, error) {
	var patient getPatientByIdResponse
	var resp Patient
	query := `query getPatientByID($id: String!) {
                getPatientByID(id: $id) {
                    id,
                    password,
                    name,
                    lastName,
                    email,
                    age,
                    height,
                    weight,
                    sex
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetDoctorById(id string) (Doctor, error) {
	var doctor getDoctorByIdResponse
	var resp Doctor
	query := `query getDoctorByID($id: String!) {
                getDoctorByID(id: $id) {
                    id,
                    password,
                    name,
                    lastName,
                    email,
					address
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func GetPatientByEmail(email string) (Patient, error) {
	var patient getPatientByEmailResponse
	var resp Patient
	query := `query getPatientByEmail($email: String!) {
                getPatientByEmail(email: $email) {
                    id,
                    password,
                    name,
                    lastName,
                    email,
                    age,
                    height,
                    weight,
                    sex
                    }
                }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetDoctorByEmail(email string) (Doctor, error) {
	var doctor getDoctorByEmailResponse
	var resp Doctor
	query := `query getDoctorByEmail($email: String!) {
                getDoctorByEmail(email: $email) {
                    id,
                    password,
                    name,
                    lastName,
                    email,
					address
                }
            }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func GetAdminByEmail(email string) (Admin, error) {
	var admin getAdminByEmailResponse
	var resp Admin
	query := `query getAdminByEmail($email: String!) {
                getAdminByEmail(email: $email) {
                    id,
                    password,
                    name,
                    lastName,
                    email
                }
            }`

	err := Query(query, map[string]interface{}{
		"email": email,
	}, &admin)
	_ = copier.Copy(&resp, &admin.Content)
	return resp, err
}

func CreatePatient(newPatient PatientInput) (Patient, error) {
	var patient createPatientResponse
	var resp Patient
	query := `mutation createPatient($email: String!, $name: String!, $lastName: String!, $password: String!, $age: Int!, $height: Int!, $weight: Int!, $sex: String!) {
            createPatient(email:$email, name:$name, lastName:$lastName, password:$password, age:$age, height:$height, weight:$weight, sex:$sex) {
                    id,
                    password,
                    name,
                    lastName,
                    email,
                    age,
                    height,
                    weight,
                    sex
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newPatient.Email,
		"name":     newPatient.Name,
		"lastName": newPatient.LastName,
		"password": newPatient.Password,
		"age":      newPatient.Age,
		"height":   newPatient.Height,
		"weight":   newPatient.Weight,
		"sex":      newPatient.Sex,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func CreateDoctor(newDoctor DoctorInput) (Doctor, error) {
	var doctor createDoctorResponse
	var resp Doctor
	query := `mutation createDoctor($email: String!, $password: String!, $name: String!, $lastName: String!, $address: String!) {
        createDoctor(email:$email, password:$password, name:$name, lastName:$lastName, address:$address) {
                    id,
                    name,
                    lastName,
                    email,
                    password,
					address
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newDoctor.Email,
		"name":     newDoctor.Name,
		"lastName": newDoctor.LastName,
		"password": newDoctor.Password,
		"address":  newDoctor.Address,
	}, &doctor)
	_ = copier.Copy(&resp, &doctor.Content)
	return resp, err
}

func CreateAdmin(newAdmin AdminInput) (Admin, error) {
	var admin createAdminResponse
	var resp Admin
	query := `mutation createAdmin($email: String!, $password: String!, $name: String!, $lastName: String!) {
        createAdmin(email:$email, password:$password, name:$name, lastName:$lastName) {
                    id,
                    name,
                    lastName,
                    email,
                    password,
                }
            }`
	err := Query(query, map[string]interface{}{
		"email":    newAdmin.Email,
		"name":     newAdmin.Name,
		"lastName": newAdmin.LastName,
		"password": newAdmin.Password,
	}, &admin)
	_ = copier.Copy(&resp, &admin.Content)
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
