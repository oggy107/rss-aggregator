package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload any) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	header := w.Header()
	header.Add("Content-Type", "application/json")

	w.WriteHeader(code)
	w.Write(jsonData)
}
