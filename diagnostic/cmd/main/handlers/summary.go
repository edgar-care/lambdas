package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	edgar_diag "github.com/edgar-care/edgarlib/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/http"
)

func GetSummary(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")

	resp := edgar_diag.GetSummary(id)
	if resp.Err != nil {
		edgarhttp.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	edgarhttp.WriteResponse(w, map[string]interface{}{
		"sessionId": resp.SessionId,
		"symptoms":  resp.Symptoms,
		"age":       resp.Age,
		"height":    resp.Height,
		"weight":    resp.Weight,
		"sex":       resp.Sex,
		"logs":      resp.Logs,
	}, 200)
}
