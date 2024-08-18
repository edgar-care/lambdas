package lib

import (
	"os"
	"net/http"
	//"fmt"

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

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) interface{} {
	_, claims, _ := jwtauth.FromContext(r.Context())
	
	//fmt.Print(claims["patient"])
	//return "hello"
	return claims["patient"].(map[string]interface{})["id"]
}
