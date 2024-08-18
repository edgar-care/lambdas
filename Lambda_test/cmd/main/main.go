package main

// import (
// 	"github.com/joho/godotenv"
// 	"github.com/ohoareau/gola"
// 	"github.com/ohoareau/gola/common"
// 	"github.com/go-chi/jwtauth/v5"
	

// 	"github.com/edgar-care/Lambda_test/cmd/main/handlers"
// 	"github.com/edgar-care/Lambda_test/cmd/main/lib"
// )

import (
	"fmt"
	"net/http"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/edgar-care/Lambda_test/cmd/main/handlers"
	"github.com/edgar-care/Lambda_test/cmd/main/lib"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Unable to load environment")
	}
}

func main() {
	addr := ":5000"
	fmt.Printf("Starting server on %v\n", addr)
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(lib.NewTokenAuth()))
	r.Get("/test/getAuthUser", handlers.AuthUser)
	return r
}

// func main() {
// 	gola.Main(common.Options{
// 		Apigw2Configurator: func(r *common.HttpRouter) {
// 			r.Use(jwtauth.Verifier(lib.NewTokenAuth()))
// 			r.Get("/test/getAuthUser", handlers.AuthUser)
// 		},
// 		Features: map[string]bool{
// 			"logger":    true,
// 			"recoverer": true,
// 			"cors":      true,
// 			"root":      true,
// 			"notfound":  true,
// 		},
// 	})
// }