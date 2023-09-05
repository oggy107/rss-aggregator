package main

import (
	"net/http"
)

type V1 struct{}

type Handler struct {
	v1 *V1
}

func (h Handler) pong(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"data": "pong"}

	respondWithJson(w, 200, data)
}

func (v1 V1) root(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"v1": "root", "age": "30"}

	respondWithJson(w, 200, data)
}

func (v1 V1) users(w http.ResponseWriter, r *http.Request) {
	// data := map[string]string{"v1": "users", "fkd": "30kk"}

	// respondWithJson(w, 200, data)
	respondWithError(w, http.StatusNotImplemented, "not implemented")
}
