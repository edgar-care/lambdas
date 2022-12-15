package handlers

import (
    "encoding/json"
    "github.com/edgar-care/auth/cmd/main/lib"
    "github.com/edgar-care/auth/cmd/main/services"
    "github.com/go-chi/chi/v5"
    "log"
    "net/http"
)


func Register(w http.ResponseWriter, req *http.Request) {
    t := chi.URLParam(req, "type")

    var token string

    if t == "d" {
        var input services.DoctorInput
        err := json.NewDecoder(req.Body).Decode(&input)
        lib.CheckError(err)

        input.Password = lib.HashPassword(input.Password)
        doctor, err := services.CreateDoctor(input)
        if err != nil {
            lib.WriteResponse(w, map[string]string{
                "message": "User already exists.",
            }, 400)
            return
        }
        token, err = lib.CreateToken(map[string]interface{}{
            "doctor": doctor,
        })
    } else {
        var input services.PatientInput

        err := json.NewDecoder(req.Body).Decode(&input)
        lib.CheckError(err)

        input.Password = lib.HashPassword(input.Password)
        patient, err := services.CreatePatient(input)
        if err != nil {
            log.Print(err.Error())
            lib.WriteResponse(w, map[string]string{
                "message": "User already exists.",
                }, 400)
            return
        }
        token, err = lib.CreateToken(map[string]interface{}{
            "patient": patient,
        })
    }

    lib.WriteResponse(w, map[string]string{
        "token": token,
    }, 200)
}