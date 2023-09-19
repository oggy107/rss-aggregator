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
func (v v1) CreateFeedFollows(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		FeedId string `json:"feed_id"`
	}

	var parameters Parameters
	decodeErr := json.NewDecoder(r.Body).Decode(&parameters)

	if decodeErr != nil {
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", decodeErr))
		return
	}

	if parameters.FeedId == "" {
		respond.WithError(w, http.StatusBadRequest, "Invalid request payload: feed_id is required")
		return
	}

	ctx := r.Context()

	user := ctx.Value(config.USER_CTX).(database.User)
	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	feedId, uuidParseErr := uuid.Parse(parameters.FeedId)

	if uuidParseErr != nil {
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", uuidParseErr))
		return
	}

	feedFollow, err := config.DB.CreateFeedFollows(ctx, database.CreateFeedFollowsParams{
		UserID: user.ID,
		FeedID: feedId,
	})

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}
