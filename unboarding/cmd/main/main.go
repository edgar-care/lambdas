package main

import (
	"context"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/unboarding/cmd/main/handlers"
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
			r.Post("/onboarding/infos", timeoutHandler(handlers.Info, 10*time.Second))
			r.Post("/onboarding/health", timeoutHandler(handlers.Health, 10*time.Second))
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

// timeoutHandler wraps the given http.HandlerFunc with a timeout duration.
func timeoutHandler(h http.HandlerFunc, timeout time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()

		r = r.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			h.ServeHTTP(w, r)
			close(done)
		}()

		select {
		case <-done:
			// Request completed within the timeout duration
		case <-ctx.Done():
			// Timeout reached
			http.Error(w, "Timeout", http.StatusGatewayTimeout)
		}
	}
}
