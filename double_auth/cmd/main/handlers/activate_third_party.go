package handlers

import (
	"encoding/json"
	"github.com/edgar-care/double_auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/pquerna/otp/totp"
	"net/http"
)

type verify_token struct {
	Token string `json:"token"`
}

func ActivateThirdParty(w http.ResponseWriter, req *http.Request) {
	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID.Code == 409 || ownerID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": ownerID.Err.Error(),
		}, ownerID.Code)
		return
	}
	if ownerID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input verify_token
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	secret := edgarlib.GetSecretThirdParty(ownerID.ID)

	valid := totp.Validate(input.Token, secret.Secret)
	if !valid {
		lib.WriteResponse(w, map[string]string{
			"message": "Token is invalid or user doesn't exist",
		}, http.StatusBadRequest)
		return
	}

	activate := edgarlib.ActivateDoubleAuthTier(ownerID.ID)
	if activate.Err != nil {
		lib.WriteError(w, activate.Code, activate.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"otp_verified": true,
	}, http.StatusOK)
}
