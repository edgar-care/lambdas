package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/unboarding/cmd/main/lib"
	"github.com/edgar-care/unboarding/cmd/main/services"
)



func Health(w http.ResponseWriter, req *http.Request) {
	var input services.HealthInput
	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	health, err := services.CreateHealth(input)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info Health (check if you share all information) " + err.Error(),
		}, 400)
		return
	}
	lib.WriteResponse(w, map[string]interface{}{
		"health": health,
	}, 200)
}