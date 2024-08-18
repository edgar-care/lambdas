package main

import (
	"github.com/edgar-care/chat/cmd/main/handlers"
	"github.com/edgar-care/chat/cmd/main/lib"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"
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
				router.Post("/ws/connection", handlers.Connection)
				router.Post("/ws/disconnect", handlers.Disconnect)
				router.Post("/ws/ready", handlers.Ready)
				router.Post("/ws/create_chat", handlers.CreatChat)
				router.Post("/ws/send_message", handlers.SendMessage)
				router.Post("/ws/get_messages", handlers.GetMessages)
				router.Post("/ws/read_message", handlers.ReadMessage)

				router.Post("/{env}/ws/connection", handlers.Connection)
				router.Post("/{env}/ws/disconnect", handlers.Disconnect)
				router.Post("/{env}/ws/ready", handlers.Ready)
				router.Post("/{env}/ws/create_chat", handlers.CreatChat)
				router.Post("/{env}/ws/send_message", handlers.SendMessage)
				router.Post("/{env}/ws/get_messages", handlers.GetMessages)
				router.Post("/{env}/ws/read_message", handlers.ReadMessage)
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
