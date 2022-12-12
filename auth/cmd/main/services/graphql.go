package services

import (
    "context"
    "github.com/machinebox/graphql"
    "os"
)

type Patient struct {
    Id          string `json:"id"`
    Password    string `json:"password"`
    Name        string `json:"name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
    Age         int    `json:"age"`
    Height      int    `json:"height"`
    Weight      int    `json:"weight"`
    Sex         string `json:"sex"`
}

type PatientInput struct {
    Password    string `json:"password"`
    Name        string `json:"name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
    Age         int    `json:"age"`
    Height      int    `json:"height"`
    Weight      int    `json:"weight"`
    Sex         string `json:"sex"`
}

type Doctor struct {
    Id          string `json:"id"`
    Password    string `json:"password"`
    Name        string `json:"name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
}

type DoctorInput struct {
    Password    string `json:"password"`
    Name        string `json:"name"`
    LastName    string `json:"last_name"`
    Email       string `json:"email"`
}

func GetPatientById(id string) (Patient, error) {
    var patient Patient
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
    return patient, err
}

func GetDoctorById(id string) (Doctor, error){
    var doctor Doctor
    query := `query getDoctorByID($id: String!) {
                getDoctorByID(id: $id) {
                    id,
                    password,
                    name,
                    lastName,
                    email
                }
            }`

    err := Query(query, map[string]interface{}{
        "id": id,
    }, &doctor)
    return doctor, err
}

func GetPatientByEmail(email string) (Patient, error) {
    var patient Patient
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
    return patient, err
}

func GetDoctorByEmail(email string) (Doctor, error) {
    var doctor Doctor
    query := `query getDoctorByEmail($email: String!) {
                getDoctorByEmail(email: $email) {
                    id,
                    password,
                    name,
                    lastName,
                    email
                }
            }`

    err := Query(query, map[string]interface{}{
        "email": email,
        }, &doctor)
    return doctor, err
}

func CreatePatient(newPatient PatientInput) (Patient, error) {
    var createdPatient Patient
    query := `mutation createPatient($email: String!, $name: String!, $lastName: String!, $password: String!, $age: Int!, $height: Int!, $weight: Int!, $sex: String!) {
                createPatient(email:$email, name:$name, password:$password, age:$age, height:$height, weight:$weight, sex:$sex) {
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
            "email": newPatient.Email,
            "name": newPatient.Name,
            "lastName": newPatient.LastName,
            "password": newPatient.Password,
            "age": newPatient.Age,
            "height": newPatient.Height,
            "weight": newPatient.Weight,
            "sex": newPatient.Sex,
            }, &createdPatient)
    return createdPatient, err
}

func CreateDoctor(newDoctor DoctorInput) (Doctor, error){
    var doctor Doctor
    query := `mutation createUsers($email: String!, $password: String!, $name: String!, $lastName: String!) {
                createUser(email:$email, password:$password, name:$name, lastName:$LastName) {
                    id,
                    name,
                    lastName,
                    email,
                    password
                }
            }`
    err := Query(query, map[string]interface{}{
        "email": newDoctor.Email,
        "name": newDoctor.Name,
        "lastName": newDoctor.LastName,
        "password": newDoctor.Password,
        }, &doctor)
    return doctor, err
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
    return createClient().Run(ctx, request, &respData)
}
