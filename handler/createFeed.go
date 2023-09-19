package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

// authorizedOnly
func (v v1) CreateFeed(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
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

	if parameters.Url == "" {
		respond.WithError(w, http.StatusBadRequest, "Invalid request payload: url is required")
		return
	}

	ctx := r.Context()

	user := ctx.Value(config.USER_CTX).(database.User)
	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	newFeed, err := config.DB.CreateFeed(ctx, database.CreateFeedParams{
		Name: parameters.Name,
		Url:  parameters.Url,
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		},
	})

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusCreated, databaseFeedtoFeed(newFeed))
}
