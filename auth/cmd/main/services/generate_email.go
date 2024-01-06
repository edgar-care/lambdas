package services

import (
	"fmt"
	"math/rand"
)

var characters_email = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func GenerateEmail(account_type string) string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = characters_email[rand.Intn(len(characters_email))]
	}
	return fmt.Sprintf("%s.%s@edgar-sante.fr", account_type, string(b))
}
