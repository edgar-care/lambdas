package lib

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/jwtauth/v5"
)

func NewTokenAuth() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	return tokenAuth
}

func VerifyToken(tokenString string) bool {
	token, err := jwtauth.VerifyToken(NewTokenAuth(), tokenString)
	if err != nil || token == nil {
		return false
	}
	return true
}

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())
	return claims["patient"].(map[string]interface{})["id"].(string)
}

func GetAuthenticatedMedecin(w http.ResponseWriter, r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())
	return claims["doctor"].(map[string]interface{})["id"].(string)
}

func AuthMiddlewareDoctor(w http.ResponseWriter, r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return ""
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return ""
	}
	return GetAuthenticatedMedecin(w, r)
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return ""
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return ""
	}
	return GetAuthenticatedUser(w, r)
}
