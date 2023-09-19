package handler

import (
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

func (v v1) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	feeds, err := config.DB.GetAllFeeds(ctx)

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
