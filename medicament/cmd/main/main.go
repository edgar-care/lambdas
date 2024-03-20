package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/medicament/cmd/main/handlers"
	"github.com/edgar-care/medicament/cmd/main/lib"
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
				router.Get("/medicament/{id}", handlers.GetMedicament)
				router.Post("/medicament", handlers.Addmedicament)
				router.Get("/medicaments", handlers.GetMedicaments)

				router.Get("/{env}/medicament/{id}", handlers.GetMedicament)
				router.Post("/{env}/medicament", handlers.Addmedicament)
				router.Get("/{env}/medicaments", handlers.GetMedicaments)
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
