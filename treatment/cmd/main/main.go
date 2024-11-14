package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/treatment/cmd/main/handlers"
	"github.com/edgar-care/treatment/cmd/main/lib"
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
				router.Get("/dashboard/treatment/{id}", handlers.GetTreatment)
				router.Post("/dashboard/treatment", handlers.Addtreatment)
				router.Get("/dashboard/treatments", handlers.GetTreatments)
				router.Put("/dashboard/treatment/{id}", handlers.EditTreatment)
				router.Delete("/dashboard/treatment/{id}", handlers.DeleteTreatment)

				router.Get("/{env}/dashboard/treatment/{id}", handlers.GetTreatment)
				router.Post("/{env}/dashboard/treatment", handlers.Addtreatment)
				router.Get("/{env}/dashboard/treatments", handlers.GetTreatments)
				router.Put("/{env}/dashboard/treatment/{id}", handlers.EditTreatment)
				router.Delete("/{env}/dashboard/treatment/{id}", handlers.DeleteTreatment)
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
