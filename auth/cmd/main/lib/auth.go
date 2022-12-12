package lib

import (
    "os"
    "strconv"

    "github.com/go-chi/jwtauth/v5"
    "golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func CreateToken(claims map[string]interface{}) (string, error) {
    _, token, err := tokenAuth.Encode(claims)
    return token, err
}

func HashPassword(password string) string {
    salt, _ := strconv.Atoi(os.Getenv("SALT"))
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
    return string(bytes)
}

func CheckPassword(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}