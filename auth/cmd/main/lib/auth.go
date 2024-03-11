package lib

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
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

	if VerifyToken(authToken) == false {
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

func GetBearerToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	//checkTokenCode, _ := edgarlib.TokenCheck(parts[1])
	//if checkTokenCode == http.StatusUnauthorized || checkTokenCode == http.StatusInternalServerError {
	//	return ""
	//}

	return parts[1]
}
