package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

type V1 struct{}

type Handler struct {
	v1 *V1
}

func (h Handler) pong(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"data": "pong"}

	respond.WithJson(w, 200, data)
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
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", decodeErr))
		return
	}

	if parameters.Name == "" {
		respond.WithError(w, http.StatusBadRequest, "Invalid request payload: name is required")
		return
	}

	ctx := r.Context()

	config := ctx.Value(CONFIG_CTX).(*ApiConfig)
	newUser, err := config.DB.CreateUser(ctx, parameters.Name)

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusCreated, databaseUsertoUser(newUser))
}

// authorizedOnly
func (v1 V1) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, ok := ctx.Value(USER_CTX).(database.User)

	if !ok {
		respond.WithError(w, http.StatusUnauthorized, "Unauthorized")
	}

	respond.WithJson(w, http.StatusOK, databaseUsertoUser(user))
}
