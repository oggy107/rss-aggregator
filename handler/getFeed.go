package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/respond"
	"github.com/oggy107/rss-aggregator/utils"
)

func (v v1) GetFeed(w http.ResponseWriter, r *http.Request) {
	rawFeedId := strings.Split(r.RequestURI, "/")[3]

	feedId, uuidParseErr := uuid.Parse(rawFeedId)

	if uuidParseErr != nil {
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", uuidParseErr))
		return
	}

	ctx := r.Context()

	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	feed, err := config.DB.GetFeed(ctx, feedId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respond.WithError(w, http.StatusNotFound, fmt.Sprintf("Feed with id %v not found", feedId))
			return
		}

		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedtoFeed(feed))
}
