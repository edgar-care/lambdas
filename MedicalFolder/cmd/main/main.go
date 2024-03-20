package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/MedicalFolder/cmd/main/handlers"
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
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
				router.Post("/dashboard/medical-info", handlers.AddMedicalInfo)
				router.Get("/dashboard/medical-info", handlers.GetMedicalInformation)
				router.Put("/dashboard/medical-info", handlers.ModifyFolderMedical)
				router.Put("/doctor/patient/{id}", handlers.ModifyMedicalInfo)

				router.Post("/{env}/dashboard/medical-info", handlers.AddMedicalInfo)
				router.Get("/{env}/dashboard/medical-info", handlers.GetMedicalInformation)
				router.Put("/{env}/dashboard/medical-info", handlers.ModifyFolderMedical)
				router.Put("/{env}/doctor/patient/{id}", handlers.ModifyMedicalInfo)
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
