package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

func NewTokenAuth() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	return tokenAuth
}

func CreateToken(claims map[string]interface{}) (string, error) {
	_, token, err := NewTokenAuth().Encode(claims)
	return token, err
}

func VerifyToken(tokenString string) bool {
	token, err := jwtauth.VerifyToken(NewTokenAuth(), tokenString)
	if err != nil || token == nil {
		return false
	}
	return true
}

func HashPassword(password string) string {
	salt, _ := strconv.Atoi(os.Getenv("SALT"))
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes)
}

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
