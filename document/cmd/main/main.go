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
				router.Get("/document/download/{id}", handlers.HandleDownload)
				router.Delete("/document/{id}", handlers.DeleteDocument)
				router.Delete("/document/favorite/{id}", handlers.RemoveFavorite)
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