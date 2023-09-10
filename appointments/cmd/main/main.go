package main

import (
	// "context"
	// "net/http"
	// "time"
	// "fmt"

	"github.com/joho/godotenv"
	//"github.com/gorilla/mux"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/edgar-care/appointments/cmd/main/handlers"
	"github.com/edgar-care/appointments/cmd/main/lib"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load environment")
	}
}

func main() {
	gola.Main(common.Options{
		Apigw2Configurator: func(r *common.HttpRouter) {
			r.Group(func(router chi.Router) {
				router.Use(jwtauth.Verifier(lib.NewTokenAuth()))
				router.Get("/doctor/{id}/appointments", handlers.GetRdvDoctor) //good
				router.Get("/patient/appointments", handlers.GetRdv) // good
				router.Post("/appointments/{id}", handlers.BookRdv) //Good
				router.Get("/patient/appointments/{id}", handlers.GetRdvPatient) //good
				router.Delete("/appointments/{id}", handlers.DeleteRdv) //good
				router.Put("/appointments/{id}", handlers.ModifRdv) //godd
			});
		},
		Features: map[string]bool{
			"logger":    true,
			"recoverer": true,
			"cors":      true,
			"root":      true,
			"notfound":  true,
		},
	})
}
