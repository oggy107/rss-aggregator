package handler

import (
	"net/http"

	"github.com/oggy107/rss-aggregator/config"
	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/respond"
)

// authorizedOnly
func (v v1) GetPostsForUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(config.USER_CTX).(database.User)
	config := ctx.Value(config.CONFIG_CTX).(*config.ApiConfig)

	posts, err := config.DB.GetPostsForUsers(ctx, database.GetPostsForUsersParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databasePostsToPosts(posts))
}
