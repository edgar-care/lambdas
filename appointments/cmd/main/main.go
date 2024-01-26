package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/appointments/cmd/main/handlers"
	"github.com/edgar-care/appointments/cmd/main/lib"
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
				// APPOINTMENT
				router.Get("/doctor/{id}/appointments", handlers.GetRdvDoctor)
				router.Get("/patient/appointments", handlers.GetRdv)
				router.Post("/appointments/{id}", handlers.BookRdv)
				router.Get("/patient/appointments/{id}", handlers.GetRdvPatient)
				router.Delete("/appointments/{id}", handlers.DeleteRdv)
				router.Put("/appointments/{id}", handlers.ModifRdv)
				router.Get("/doctor/appointments/{id}", handlers.GetDoctorAppointment)
				router.Get("/doctor/appointments", handlers.GetAllDoctorAppointments)
				router.Put("/doctor/appointments/{id}", handlers.UpdateDoctorAppointment)
				router.Post("/doctor/appointments", handlers.CreateRdv)
				router.Delete("/doctor/appointments/{id}", handlers.CancelRdv)

				// SLOT
				router.Get("/doctor/slot/{id}", handlers.GetSlotId)
				router.Post("/doctor/slot", handlers.CreateSlot)
				router.Delete("/doctor/slot/{id}", handlers.DeleteSlot)
				router.Get("/doctor/slots", handlers.GetSlots)
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
