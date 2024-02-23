package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/pushnotification/cmd/main/handlers"
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
			r.Post("/push-notif", timeoutHandler(handlers.Notification, 10*time.Second))
			r.Post("/{env}/push-notif", timeoutHandler(handlers.Notification, 10*time.Second))
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

func timeoutHandler(next http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		done := make(chan bool)

		// Exécute la fonction de traitement de la requête dans une goroutine
		go func() {
			next(w, r)
			done <- true
		}()

		select {
		case <-ctx.Done():
			// Le délai imparti est écoulé, renvoie une réponse d'erreur
			http.Error(w, fmt.Sprintf("Temps d'exécution de la requête dépassé (limite : %s)", timeout), http.StatusRequestTimeout)
		case <-done:
			// La fonction de traitement de la requête a terminé avant l'expiration du délai
			return
		}
	}
}
