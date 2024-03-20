package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/document/cmd/main/handlers"
	"github.com/edgar-care/document/cmd/main/lib"
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
				router.Post("/document/upload", handlers.HandleUpload)
				router.Post("/document/favorite/{id}", handlers.HandleFavorite)
				router.Post("/doctor/document/upload", handlers.UploadFromDoctor)
				router.Get("/document/download/{id}", handlers.HandleDownload)
				router.Get("/doctor/document/{id}", handlers.DownloadFromDoctor)
				router.Delete("/document/{id}", handlers.DeleteDocument)
				router.Delete("/document/favorite/{id}", handlers.RemoveFavorite)
				router.Put("/document/{id}", handlers.HandleUpdate)
				router.Get("/document/download", handlers.GetAllDocument)

				router.Post("/{env}/document/upload", handlers.HandleUpload)
				router.Post("/{env}/document/favorite/{id}", handlers.HandleFavorite)
				router.Post("/{env}/doctor/document/upload", handlers.UploadFromDoctor)
				router.Get("/{env}/document/download/{id}", handlers.HandleDownload)
				router.Get("/{env}/doctor/document/{id}", handlers.DownloadFromDoctor)
				router.Delete("/{env}/document/{id}", handlers.DeleteDocument)
				router.Delete("/{env}/document/favorite/{id}", handlers.RemoveFavorite)
				router.Put("/{env}/document/{id}", handlers.HandleUpdate)
				router.Get("/{env}/document/download", handlers.GetAllDocument)
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
