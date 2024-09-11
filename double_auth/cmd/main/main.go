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
				//Activation / desactivation 2fa method
				router.Post("/2fa/method/email", handlers.AddDoubleAutEmail)
				router.Post("/2fa/method/third_party/generate", handlers.AddDoubleAuthAppTier)
				router.Post("/2fa/method/third_party", handlers.ActivateThirdParty)
				router.Post("/2fa/method/mobile", handlers.AddDoubleMobile)
				router.Delete("/dashboard/2fa/{ENUM}", handlers.DisableDoubleAuth)
				router.Get("/dashboard/2fa", handlers.GetDoubleAuth)

				// trust device
				router.Post("/dashboard/2fa/device/{id}", handlers.AddTrustDevice)
				router.Get("/dashboard/2fa/device/{id}", handlers.GetTrustDevice)
				router.Get("/dashboard/2fa/devices", handlers.GetTrustDevices)
				router.Delete("/dashboard/2fa/device/{id}", handlers.DeleteTrustDevice)

				// device connect
				router.Delete("/dashboard/device/{id}", handlers.DeleteDevice)
				router.Get("/dashboard/device/{id}", handlers.GetDevice)
				router.Get("/dashboard/devices", handlers.GetDevices)

				//ENV Activation / desactivation 2fa method
				router.Post("/{env}/2fa/method/email", handlers.AddDoubleAutEmail)
				router.Post("/{env}/2fa/method/third_party/generate", handlers.AddDoubleAuthAppTier)
				router.Post("/{env}/2fa/method/third_party", handlers.ActivateThirdParty)
				router.Post("/{env}/2fa/method/mobile", handlers.AddDoubleMobile)
				router.Delete("/{env}/dashboard/2fa/{ENUM}", handlers.DisableDoubleAuth)
				router.Get("/{env}/dashboard/2fa", handlers.GetDoubleAuth)

				//ENV trust device
				router.Post("/{env}/dashboard/2fa/device/{id}", handlers.AddTrustDevice)
				router.Get("/{env}/dashboard/2fa/device/{id}", handlers.GetTrustDevice)
				router.Get("/{env}/dashboard/2fa/devices", handlers.GetTrustDevices)
				router.Delete("/{env}/dashboard/2fa/device/{id}", handlers.DeleteTrustDevice)

				//ENV device connect
				router.Delete("/{env}/dashboard/device/{id}", handlers.DeleteDevice)
				router.Get("/{env}/dashboard/device/{id}", handlers.GetDevice)
				router.Get("/{env}/dashboard/devices", handlers.GetDevices)

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
