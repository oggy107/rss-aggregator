package respond

import (
	"encoding/json"
	"net/http"

	"github.com/oggy107/rss-aggregator/utils"
)

func WithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		utils.LogNonFatal("Responding with 5XX status: %s", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	WithJson(w, code, ErrorResponse{Error: message})
}

func WithJson(w http.ResponseWriter, code int, payload any) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		utils.LogNonFatal("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(jsonData)
}
