package main

import (
	"github.com/joho/godotenv"
	"github.com/ohoareau/gola"
	"github.com/ohoareau/gola/common"

	"github.com/edgar-care/auth/cmd/main/handlers"
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
			r.Post("/{env}/auth/{type}/login", handlers.Login)
			r.Post("/{env}/auth/{type}/register", handlers.Register)
			r.Post("/{env}/auth/p/create_account", handlers.CreatePatientAccount)
			r.Post("/{env}/auth/{type}/missing-password", handlers.MissingPassword)
			r.Post("/{env}/auth/{type}/reset-password", handlers.ResetPassword)
			r.Post("/{env}/auth/update_password", handlers.UpdatePassword)

			r.Put("/{env}/auth/disable_account", handlers.DisableAccount)
			r.Put("/{env}/auth/enable_account", handlers.EnableAccount)
			r.Post("/{env}/auth/delete_account", handlers.DeleteAccountInit)

			r.Post("/{env}/auth/creation_backup_code", handlers.CreateBackupCode)
			r.Post("/{env}/auth/sending_email", handlers.SenderEmail2FA)
			r.Post("/{env}/auth/{type}/email_2fa", handlers.Login2faEmail)
			r.Post("/{env}/auth/{type}/backup_code_2fa", handlers.Login2faBackupCode)
			r.Post("/{env}/auth/{type}/third_party_2fa", handlers.Login2faThirdParty)
			r.Post("/{env}/auth/{type}/mobile_2fa", handlers.Login2faEmail)

			r.Post("/{env}/auth/ws/ready", handlers.ReadyLoginWeb)
			r.Post("/{env}/auth/ws/ask_mobile_connection", handlers.AskMobileConnection)
			r.Post("/{env}/auth/ws/response_mobile_connection", handlers.ResponseMobileConnection)

			r.Post("/auth/{type}/login", handlers.Login)
			r.Post("/auth/{type}/register", handlers.Register)
			r.Post("/auth/p/create_account", handlers.CreatePatientAccount)
			r.Post("/auth/{type}/missing-password", handlers.MissingPassword)
			r.Post("/auth/{type}/reset-password", handlers.ResetPassword)
			r.Post("/auth/update_password", handlers.UpdatePassword)

			r.Post("/auth/sending_email", handlers.SenderEmail2FA)
			r.Post("/auth/{type}/email_2fa", handlers.Login2faEmail)
			r.Post("/auth/{type}/backup_code_2fa", handlers.Login2faBackupCode)
			r.Post("/auth/{type}/third_party_2fa", handlers.Login2faThirdParty)
			r.Post("/auth/{type}/mobile_2fa", handlers.Login2faEmail)

			r.Put("/auth/disable_account", handlers.DisableAccount)
			r.Put("/auth/enable_account", handlers.EnableAccount)
			r.Post("/auth/delete_account", handlers.DeleteAccountInit)
			r.Post("/auth/creation_backup_code", handlers.CreateBackupCode)

			r.Post("/auth/ws/ready", handlers.ReadyLoginWeb)
			r.Post("/auth/ws/ask_mobile_connection", handlers.AskMobileConnection)
			r.Post("/auth/ws/response_mobile_connection", handlers.ResponseMobileConnection)

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
