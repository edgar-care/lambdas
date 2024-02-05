package handlers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/diagnostic/cmd/main/services"
	edgarhttp "github.com/edgar-care/edgarlib/http"
)

func GetSummary(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	session, err := services.GetSessionById(id)
	if err != nil {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": "Unable to get session: " + strings.ToLower(err.Error()[9:]),
		}, 400)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"sessionId": session.Id,
		"symptoms":  services.StringToSymptoms(session.Symptoms),
		"age":       session.Age,
		"height":    session.Height,
		"weight":    session.Weight,
		"sex":       session.Sex,
		"logs":      session.Logs,
	}, 200)
}
