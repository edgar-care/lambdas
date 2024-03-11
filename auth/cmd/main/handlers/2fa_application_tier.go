package handlers

//import (
//	"encoding/json"
//	"net/http"
//
//	"github.com/pquerna/otp/totp"
//)
//
//func DoubleAutTier(w http.ResponseWriter, req *http.Request) {
//
//	var payload models.OTPInput
//
//	// Décoder le payload JSON
//	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	message := "Token is invalid or user doesn't exist"
//
//	// Validation du token OTP
//	valid := totp.Validate(payload.Token, payload.OtpSecret)
//	if !valid {
//		http.Error(w, message, http.StatusBadRequest)
//		return
//	}
//
//	// Appeler la fonction pour sauvegarder les informations dans MongoDB
//	if err := third_application(payload); err != nil {
//		http.Error(w, "Failed to save data to MongoDB", http.StatusInternalServerError)
//		return
//	}
//
//	// Répondre avec succès
//	response := map[string]bool{"otp_valid": true}
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(response)
//}
