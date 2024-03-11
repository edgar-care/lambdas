package lib

import (
	"encoding/base64"
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"os"
	"strings"
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

func GetAuthenticatedUser(authToken string) string {
	decodedBytes, err := base64.RawURLEncoding.DecodeString(strings.Split(authToken, ".")[1])
	if err != nil {
		CheckError(err)
	}

	var jsonMap map[string]interface{}
	json.Unmarshal(decodedBytes, &jsonMap)
	if jsonMap["patient"] != nil {
		return jsonMap["patient"].(map[string]interface{})["id"].(string)
	}
	return ""
}

func GetAuthenticatedMedecin(authToken string) string {

	decodedBytes, err := base64.RawURLEncoding.DecodeString(strings.Split(authToken, ".")[1])
	if err != nil {
		CheckError(err)
	}

	var jsonMap map[string]interface{}
	json.Unmarshal(decodedBytes, &jsonMap)
	if jsonMap["doctor"] != nil {
		return jsonMap["doctor"].(map[string]interface{})["id"].(string)
	}
	return ""
}

func AuthMiddlewareDoctor(authToken string) string {
	if authToken == "" {
		return ""
	}

	if VerifyToken(authToken) == false {
		return ""
	}
	return GetAuthenticatedMedecin(authToken)
}

func AuthMiddleware(authToken string) string {
	if authToken == "" {
		return ""
	}

	if VerifyToken(authToken) == false {
		return ""
	}
	return GetAuthenticatedUser(authToken)
}
