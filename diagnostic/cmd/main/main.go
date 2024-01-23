package main

import (
	"github.com/edgar-care/diagnostic/cmd/main/handlers"
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
			r.Post("/diagnostic/initiate", handlers.Initiate)
			r.Post("/diagnostic/diagnose", handlers.Diagnose)
			r.Get("/diagnostic/summary/{id}", handlers.GetSummary)
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
