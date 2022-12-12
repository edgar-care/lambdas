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
        input.Password = lib.HashPassword(input.Password)
        err := json.NewDecoder(req.Body).Decode(&input)
        lib.CheckError(err)

        user, err := services.CreateDoctor(input)
        if err != nil {
            lib.WriteResponse(w, map[string]string{
                "message": "User already exists.",
            }, 400)
        }
        token, err = lib.CreateToken(map[string]interface{}{
            "user": user,
        })
    } else {
        var input services.PatientInput
        input.Password = lib.HashPassword(input.Password)
        err := json.NewDecoder(req.Body).Decode(&input)
        lib.CheckError(err)

        user, err := services.CreatePatient(input)
        if err != nil {
            log.Print(err.Error())
            lib.WriteResponse(w, map[string]string{
                "message": "User already exists.",
                }, 400)
            return
        }
        token, err = lib.CreateToken(map[string]interface{}{
            "user": user,
        })
    }

    lib.WriteResponse(w, map[string]string{
        "token": token,
    }, 200)
}