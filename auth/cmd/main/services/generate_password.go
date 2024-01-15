package services

import "math/rand"

var characters_password = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789*${!:;?+-&}")

func GeneratePassword(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = characters_password[rand.Intn(len(characters_password))]
	}
	return string(b)
}
