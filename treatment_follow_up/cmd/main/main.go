package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/treatment_follow_up/cmd/main/handlers"
	"github.com/edgar-care/treatment_follow_up/cmd/main/lib"
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
				router.Post("/dashboard/treatment/follow-up", handlers.AddFollowTreatment)
				router.Delete("/dashboard/treatment/follow-up/{id}", handlers.DeleteFollowTreatment)
				router.Get("/dashboard/treatment/follow-up/{id}", handlers.GetFollowTreatment)
				router.Get("/dashboard/treatment/follow-up", handlers.GetfFollowsTreatments)

				router.Post("/{env}/dashboard/treatment/follow-up", handlers.AddFollowTreatment)
				router.Delete("/{env}/dashboard/treatment/follow-up/{id}", handlers.DeleteFollowTreatment)
				router.Get("/{env}/dashboard/treatment/follow-up/{id}", handlers.GetFollowTreatment)
				router.Get("/{env}/dashboard/treatment/follow-up", handlers.GetfFollowsTreatments)
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
