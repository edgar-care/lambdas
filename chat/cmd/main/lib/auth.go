package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

//func GetAuthenticatedUser(authToken string) string {
//	decodedBytes, err := base64.RawURLEncoding.DecodeString(strings.Split(authToken, ".")[1])
//	if err != nil {
//		CheckError(err)
//	}
//
//	var jsonMap map[string]interface{}
//	json.Unmarshal(decodedBytes, &jsonMap)
//	if jsonMap["patient"] != nil {
//		return jsonMap["patient"].(map[string]interface{})["id"].(string)
//	}
//	return ""
//}
//
//func GetAuthenticatedMedecin(authToken string) string {
//
//	decodedBytes, err := base64.RawURLEncoding.DecodeString(strings.Split(authToken, ".")[1])
//	if err != nil {
//		CheckError(err)
//	}
//
//	var jsonMap map[string]interface{}
//	json.Unmarshal(decodedBytes, &jsonMap)
//	if jsonMap["doctor"] != nil {
//		return jsonMap["doctor"].(map[string]interface{})["id"].(string)
//	}
//	return ""
//}
//
//func AuthMiddlewareDoctor(authToken string) string {
//	if authToken == "" {
//		return ""
//	}
//
//	if VerifyToken(authToken) == false {
//		return ""
//	}
//	return GetAuthenticatedMedecin(authToken)
//}
//
//func AuthMiddleware(authToken string) string {
//	if authToken == "" {
//		return ""
//	}
//
//	if VerifyToken(authToken) == false {
//		return ""
//	}
//	return GetAuthenticatedUser(authToken)
//}

func AuthMiddleware(authToken string) string {
	if authToken == "" {
		return ""
	}

	if !VerifyToken(authToken) {
		return ""
	}
	return GetAuthenticatedUser(authToken)
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

	// Vérifie si l'utilisateur est un patient
	if patientID, ok := jsonMap["id"].(string); ok {
		if patientEmail, ok := jsonMap["patient"].(string); ok {
			fmt.Printf("Patient Email: %s, ID: %s\n", patientEmail, patientID)
			return patientID
		}
	}

	// Vérifie si l'utilisateur est un docteur
	if doctorID, ok := jsonMap["id"].(string); ok {
		if doctorEmail, ok := jsonMap["doctor"].(string); ok {
			fmt.Printf("Doctor Email: %s, ID: %s\n", doctorEmail, doctorID)
			return doctorID
		}
	}

	return ""
}

func GetDeviceId(authToken string) string {
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

	deviceID, _ := jsonMap["name_device"].(string)

	return deviceID
}
