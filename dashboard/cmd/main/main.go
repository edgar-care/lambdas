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
	// c := cron.New()
	// c.AddFunc("0 0 * * *", func() {
	// 	services.CronJobDeleteExpiredAccounts()
	// })
	// c.Start()

	// // Démarrage du cron job pour l'envoi quotidien des emails d'expiration
	// c.AddFunc("0 0 * * *", func() {
	// 	services.CronJobSendExpirationEmails()
	// })
	// c.Start()
	gola.Main(common.Options{
		Apigw2Configurator: func(r *common.HttpRouter) {
			r.Group(func(router chi.Router) {
				router.Use(jwtauth.Verifier(lib.NewTokenAuth()))
				// Get
				router.Get("/doctor/patient/{id}", handlers.GetPatientId)
				router.Get("/doctor/patients", handlers.GetPatients)
				// router.Post("/doctor/patient", handlers.CreatePatient)
				// router.Put("/doctor/patient/{id}", handlers.UpdatePatientInfo)
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
