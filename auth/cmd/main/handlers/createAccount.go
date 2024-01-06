package handlers

import (
	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

func CreateAccount(w http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")

	if req.Header.Get("admin_token") != os.Getenv("ADMIN_TOKEN") {
		lib.WriteResponse(w, map[string]string{}, 401)
		return
	}

	var token string
	var email string
	var password string

	if t == "test" {
		var input services.TestAccountInput
		email = services.GenerateEmail("test")
		password = services.GeneratePassword(10)

		input.Email = email
		input.Password = lib.HashPassword(password)

		test, err := services.CreateTestAccount(input)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create test account: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		token, _ = lib.CreateToken(map[string]interface{}{
			"test": test,
		})
	} else if t == "demo" {
		var input services.DemoAccountInput
		email = services.GenerateEmail("demo")
		password = services.GeneratePassword(10)

		input.Email = email
		input.Password = lib.HashPassword(password)

		demo, err := services.CreateDemoAccount(input)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create demo account: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		token, _ = lib.CreateToken(map[string]interface{}{
			"demo": demo,
		})
	} else {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid type",
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]string{
		"email":      email,
		"password":   password,
		"auth_token": token,
	}, 200)
}
