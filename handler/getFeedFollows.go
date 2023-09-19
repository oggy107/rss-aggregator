package handler

import (
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

// authorizedOnly
func (v v1) GetFeedFollows(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(config.USER_CTX).(database.User)
	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	feedFollows, err := config.DB.GetFeedFollows(ctx, user.ID)

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}
