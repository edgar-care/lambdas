package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/unboarding/cmd/main/lib"
	"github.com/edgar-care/unboarding/cmd/main/services"
)

func Info(w http.ResponseWriter, req *http.Request) {
	var input services.InfoInput
	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	info, err := services.CreateInfo(input)
	//info, err := services.GetInfoById(input.Id)
	
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info (check if you share all information) " + err.Error(),
		}, 400)
		return
	}
	lib.WriteResponse(w, map[string]interface{}{
		"info": info,
	}, 200)
}