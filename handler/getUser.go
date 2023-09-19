package handler

import (
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
)

// authorizedOnly
func (v v1) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(config.USER_CTX).(database.User)

	respond.WithJson(w, http.StatusOK, databaseUsertoUser(user))
}
