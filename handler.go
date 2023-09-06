package main

import (
	"encoding/json"
	"fmt"
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

// func (v1 V1) root(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]string{"v1": "root", "age": "30"}

// 	respondWithJson(w, 200, data)
// }

func (v1 V1) createUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
	}

	var parameters Parameters
	decodeErr := json.NewDecoder(r.Body).Decode(&parameters)

	if decodeErr != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", decodeErr))
		return
	}

	if parameters.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload: name is required")
		return
	}

	ctx := r.Context()

	config := ctx.Value("config").(*ApiConfig)
	newUser, err := config.DB.CreateUser(ctx, parameters.Name)

	if err != nil {
		logNonFatal(err.Error())
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithJson(w, 200, databaseUsertoUser(newUser))
}
