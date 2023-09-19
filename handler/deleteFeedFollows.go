package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

// authorizedOnly
func (v v1) DeleteFeedFollows(w http.ResponseWriter, r *http.Request) {
	rawFeedId := chi.URLParam(r, "feed_id")

	feedId, uuidParseErr := uuid.Parse(rawFeedId)

	if uuidParseErr != nil {
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", uuidParseErr))
		return
	}

	ctx := r.Context()

	apiConfig := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)
	user := ctx.Value(config.USER_CTX).(database.User)

	err := apiConfig.DB.DeleteFeedFollows(ctx, database.DeleteFeedFollowsParams{
		ID:     feedId,
		UserID: user.ID,
	})

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, map[string]bool{"success": true})
}
