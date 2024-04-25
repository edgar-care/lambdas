package main

import (
	"github.com/edgar-care/diagnostic/cmd/main/handlers"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"
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

				router.Post("/diagnostic/initiate", handlers.Initiate)
				router.Post("/diagnostic/diagnose", handlers.Diagnose)
				router.Get("/diagnostic/summary/{id}", handlers.GetSummary)

				router.Post("/{env}/diagnostic/initiate", handlers.Initiate)
				router.Post("/{env}/diagnostic/diagnose", handlers.Diagnose)
				router.Get("/{env}/diagnostic/summary/{id}", handlers.GetSummary)
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
