package main

import (
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/auth/cmd/main/handlers"
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
			r.Post("/{env}/auth/{type}/login", handlers.Login)
			r.Post("/{env}/auth/{type}/register", handlers.Register)
			r.Post("/{env}/auth/p/create_account", handlers.CreatePatientAccount)
			r.Post("/{env}/auth/p/missing-password", handlers.MissingPassword)
			r.Post("/{env}/auth/p/reset-password", handlers.ResetPassword)

			r.Post("/auth/{type}/login", handlers.Login)
			r.Post("/auth/{type}/register", handlers.Register)
			r.Post("/auth/p/create_account", handlers.CreatePatientAccount)
			r.Post("/auth/p/missing-password", handlers.MissingPassword)
			r.Post("/auth/p/reset-password", handlers.ResetPassword)
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
