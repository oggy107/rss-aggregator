package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/oggy107/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Name      string        `json:"name"`
	Url       string        `json:"url"`
	UserID    uuid.NullUUID `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) Post {
	var description *string

	if post.Description.Valid {
		description = &post.Description.String
	}

	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Description: description,
		PublishedAt: post.PublishedAt,
		Url:         post.Url,
		FeedID:      post.FeedID,
	}
}

func databasePostsToPosts(posts []database.Post) []Post {
	parsedPosts := []Post{}

	for _, post := range posts {
		parsedPosts = append(parsedPosts, databasePostToPost(post))
	}

	return parsedPosts
}

func databaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedId:    feedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	parsedFeedFollows := []FeedFollow{}

	for _, feedFollow := range feedFollows {
		parsedFeedFollows = append(parsedFeedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}

	return parsedFeedFollows
}

func databaseUsertoUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func databaseFeedtoFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Name,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	parsedFeeds := []Feed{}

	for _, feed := range feeds {
		parsedFeeds = append(parsedFeeds, databaseFeedtoFeed(feed))
	}

	return parsedFeeds
}
