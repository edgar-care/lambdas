package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

//func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) string {
//	_, claims, _ := jwtauth.FromContext(r.Context())
//	return claims["patient"].(map[string]interface{})["id"].(string)
//}

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
	return GetAuthenticatedUser(reqToken)
}

func GetAuthenticatedUser(authToken string) string {
	parts := strings.Split(authToken, ".")
	if len(parts) != 3 {
		return ""
	}
	decodedBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		CheckError(err)
		return ""
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(decodedBytes, &jsonMap); err != nil {
		CheckError(err)
		return ""
	}

	if id, ok := jsonMap["id"].(string); ok {
		return id
	}

	if patient, ok := jsonMap["patient"].(map[string]interface{}); ok {
		if id, ok := patient["id"].(string); ok {
			fmt.Print(id)
			return id
		}
	}

	return ""
}

//func GetAuthenticated(authToken string) string {
//	parts := strings.Split(authToken, ".")
//	if len(parts) != 3 {
//		return ""
//	}
//	decodedBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
//	if err != nil {
//		CheckError(err)
//		return ""
//	}
//
//	var jsonMap map[string]interface{}
//	if err := json.Unmarshal(decodedBytes, &jsonMap); err != nil {
//		CheckError(err)
//		return ""
//	}
//
//	if id, ok := jsonMap["id"].(string); ok {
//		return id
//	}
//
//	if patient, ok := jsonMap["patient"].(map[string]interface{}); ok {
//		if id, ok := patient["id"].(string); ok {
//			fmt.Print(id)
//			return id
//		}
//	}
//
//	if doctor, ok := jsonMap["doctor"].(map[string]interface{}); ok {
//		if id, ok := doctor["id"].(string); ok {
//			fmt.Print(id)
//			return id
//		}
//	}
//
//	return ""
//}
