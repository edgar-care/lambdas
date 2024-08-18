package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	edgar_diag "github.com/edgar-care/edgarlib/v2/diagnostic"
	edgarhttp "github.com/edgar-care/edgarlib/v2/http"
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
		"session_id": resp.SessionId,
		"diseases":   resp.Diseases,
		"fiability":  resp.Fiability,
		"symptoms":   resp.Symptoms,
		"logs":       resp.Logs,
		"alerts":     resp.Alerts,
	}, 200)
}
