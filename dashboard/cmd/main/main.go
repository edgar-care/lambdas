package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/dashboard/cmd/main/handlers"
	"github.com/edgar-care/dashboard/cmd/main/lib"
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
				router.Get("/doctor/patient/{id}", handlers.GetPatientId)
				router.Get("/doctor/patients", handlers.GetPatients)
				router.Post("/doctor/patient", handlers.CreatePatient)
				router.Delete("/doctor/patient/{id}", handlers.DeletePatientHandler)
				router.Post("/doctor/diagnostic/{id}", handlers.RevPreDiagnostic)
				router.Get("/doctor/diagnostic/waiting", handlers.GetPreDignosticWait)
				router.Get("/doctor/{id}", handlers.GetDoctorId)
				router.Get("/doctors", handlers.GetDoctors)
			})
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
