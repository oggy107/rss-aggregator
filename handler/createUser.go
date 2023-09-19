package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

func (v v1) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)
	newUser, err := config.DB.CreateUser(ctx, parameters.Name)

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusCreated, databaseUsertoUser(newUser))
}
