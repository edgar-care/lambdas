package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/double_auth/cmd/main/handlers"
	"github.com/edgar-care/double_auth/cmd/main/lib"
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
				router.Get("/dashboard/devices", handlers.GetDevices)
				router.Get("/dashboard/double_auth/{id}", handlers.GetDevice)
				router.Delete("/dashboard/double_auth/{id}", handlers.DeleteDevice)
				router.Post("/dashboard/double_auth/{id}", handlers.DeleteDevice)
				router.Post("/2fa/method/email", handlers.AddDoubleAutEmail)
				router.Post("/2fa/method/app-tier", handlers.AddDoubleAuthAppTier)
				router.Post("/2fa/method/mobile", handlers.AddDoubleMobile)
				router.Delete("/2fa/method/{ENUM}", handlers.DisableDoubleAuth)

				router.Post("/dashboard/2fa/device/{id}", handlers.AddTrustDevice)
				router.Get("/dashboard/2fa/device/{id}", handlers.GetTrustDevice)
				router.Get("/dashboard/2fa/devices", handlers.GetDevices)
				router.Delete("/dashboard/2fa/device/{id}", handlers.DeleteTrustDevice)

				router.Get("/{env}/dashboard/devices", handlers.GetDevices)
				router.Get("/{env}/dashboard/double_auth/{id}", handlers.GetDevice)
				router.Delete("/{env}/dashboard/double_auth/{id}", handlers.DeleteDevice)
				router.Post("/{env}/dashboard/double_auth/{id}", handlers.DeleteDevice)
				router.Post("/{env}/2fa/method/email", handlers.AddDoubleAutEmail)
				router.Post("/{env}/2fa/method/app-tier", handlers.AddDoubleAuthAppTier)
				router.Post("/{env}/2fa/method/mobile", handlers.AddDoubleMobile)
				router.Delete("/{env}/2fa/method/{ENUM}", handlers.DisableDoubleAuth)

				router.Post("/{env}/dashboard/2fa/device/{id}", handlers.AddTrustDevice)
				router.Get("/{env}/dashboard/2fa/device/{id}", handlers.GetTrustDevice)
				router.Get("/{env}/dashboard/2fa/devices", handlers.GetDevices)
				router.Delete("/{env}/dashboard/2fa/device/{id}", handlers.DeleteTrustDevice)
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
