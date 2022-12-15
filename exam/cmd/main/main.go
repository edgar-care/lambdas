package main

import (
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/exam/cmd/main/handlers"
)

func main() {
	gola.Main(common.Options{
		Apigw2Configurator: func(r *common.HttpRouter) {
			r.Post("/exam", handlers.Exam)
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
