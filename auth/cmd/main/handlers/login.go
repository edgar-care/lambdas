package handlers

import (
    "encoding/json"
    //"github.com/davecgh/go-spew/spew"
    "github.com/edgar-care/auth/cmd/main/lib"
    "github.com/edgar-care/auth/cmd/main/services"
    "github.com/go-chi/chi/v5"
    "net/http"
)

type loginInput struct {
    Email       string `json:"email"`
    Password    string `json:"password"`
}

func Login(w http.ResponseWriter, req *http.Request) {
    var input loginInput
    var user interface{}
    t := chi.URLParam(req, "type")

    err := json.NewDecoder(req.Body).Decode(&input)
    lib.CheckError(err)

    if t == "d" {
        user, err = services.GetDoctorByEmail(input.Email)
    } else {
        user, err = services.GetPatientByEmail(input.Email)
    }

    if !(t == "d" && lib.CheckPassword(input.Password, user.(services.Doctor).Password)) &&
        !(t == "p" && lib.CheckPassword(input.Password, user.(services.Patient).Password)){
        lib.WriteResponse(w, map[string]string{
            "message": "Username and password mismatch.",
            }, 400)
        return
    }

    token, err := lib.CreateToken(map[string]interface{}{
        "user": user,
    })

    lib.CheckError(err)
    lib.WriteResponse(w, map[string]interface{}{
        "token": token,
    }, 200)
}