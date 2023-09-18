package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
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

	user := ctx.Value(USER_CTX).(database.User)

	respond.WithJson(w, http.StatusOK, databaseUsertoUser(user))
}

// authorizedOnly
func (v1 V1) createFeed(w http.ResponseWriter, r *http.Request) {
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

	user := ctx.Value(USER_CTX).(database.User)
	config := ctx.Value(CONFIG_CTX).(*ApiConfig)

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

func (v1 V1) getFeed(w http.ResponseWriter, r *http.Request) {
	rawFeedId := strings.Split(r.RequestURI, "/")[3]

	feedId, uuidParseErr := uuid.Parse(rawFeedId)

	if uuidParseErr != nil {
		respond.WithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", uuidParseErr))
		return
	}

	ctx := r.Context()

	config := ctx.Value(CONFIG_CTX).(*ApiConfig)

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

// authorizedOnly
func (v1 V1) getFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := ctx.Value(USER_CTX).(database.User)
	config := ctx.Value(CONFIG_CTX).(*ApiConfig)

	feeds, err := config.DB.GetFeeds(ctx, uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	})

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

func (v1 V1) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	config := ctx.Value(CONFIG_CTX).(*ApiConfig)

	feeds, err := config.DB.GetAllFeeds(ctx)

	if err != nil {
		utils.LogNonFatal(err.Error())
		respond.WithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respond.WithJson(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}

// authorizedOnly
func (v1 V1) CreateFeedFollows(w http.ResponseWriter, r *http.Request) {
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

	user := ctx.Value(USER_CTX).(database.User)
	config := ctx.Value(CONFIG_CTX).(*ApiConfig)

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
