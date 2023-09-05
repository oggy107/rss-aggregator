package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func logError(format string, a ...any) {
	log.Print("[ERROR]: ", fmt.Sprintf(format, a...))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		logError("Responding with 5XX status: %s", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, ErrorResponse{Error: message})
}

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		logError("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(jsonData)
}
