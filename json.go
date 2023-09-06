package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		logNonFatal("Responding with 5XX status: %s", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, ErrorResponse{Error: message})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		logNonFatal("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(jsonData)
}
