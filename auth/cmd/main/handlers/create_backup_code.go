package handlers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"math/big"
	"net/http"
	"strings"

	"github.com/edgar-care/auth/cmd/main/lib"
)

type CreateSaveCodeResponse struct {
	SaveCode model.SaveCode
	Code     int
	Err      error
}

func CreateBackupCode(w http.ResponseWriter, req *http.Request) {

	authToken := lib.GetBearerToken(req)
	patientID := lib.AuthMiddleware(authToken)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	backupCode := CreateBackupCodes(patientID, authToken, req)

	if backupCode.Err != nil {
		lib.WriteError(w, backupCode.Code, backupCode.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": backupCode.SaveCode,
	}, backupCode.Code)
}

func generateRandomCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func generateBackupCodes() ([]string, error) {
	const codeLength = 8
	const numCodes = 10
	codes := make([]string, numCodes)

	for i := 0; i < numCodes; i++ {
		code, err := generateRandomCode(codeLength)
		if err != nil {
			return nil, err
		}
		codes[i] = code
	}

	return codes, nil
}

func hashCode(code string) string {
	hash := sha256.Sum256([]byte(code))
	return hex.EncodeToString(hash[:])
}

func CreateBackupCodes(id string, token string, req *http.Request) CreateSaveCodeResponse {
	gqlClient := graphql.CreateClient()

	codes, err := generateBackupCodes()
	if err != nil {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
	}

	hashedCodes := make([]string, len(codes))
	for i, code := range codes {
		hashedCodes[i] = hashCode(code)
	}

	saveCode, err := graphql.CreateSaveCode(context.Background(), gqlClient, hashedCodes)
	if err != nil {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 400, Err: errors.New("unable to create save code")}
	}

	accountType, err := GetAccountType(token)
	if accountType == "" {
		return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 400, Err: errors.New("no account type found")}
	}

	if accountType == "patient" {
		patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient")}
		}
		doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a double-auth")}
		}

		_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, doubleAuth.GetDoubleAuthById.Methods, saveCode.CreateSaveCode.Id, doubleAuth.GetDoubleAuthById.Url, doubleAuth.GetDoubleAuthById.Trust_device_id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
		}

	} else {
		doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a doctor")}
		}
		doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, doctor.GetDoctorById.Double_auth_methods_id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a double-auth")}
		}

		_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, doubleAuth.GetDoubleAuthById.Methods, saveCode.CreateSaveCode.Id, doubleAuth.GetDoubleAuthById.Url, doubleAuth.GetDoubleAuthById.Trust_device_id)
		if err != nil {
			return CreateSaveCodeResponse{SaveCode: model.SaveCode{}, Code: 500, Err: err}
		}
	}

	return CreateSaveCodeResponse{
		SaveCode: model.SaveCode{
			Code: codes,
		},
		Code: 201,
		Err:  nil,
	}
}

func GetAccountType(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", err
	}
	if _, valid := claims["patient"].(string); valid {
		return "patient", nil
	}
	if _, valid := claims["doctor"].(string); valid {
		return "doctor", nil
	}

	return "", errors.New("no account type found")
}
