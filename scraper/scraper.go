package scraper

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/oggy107/rss-aggregator/internal/database"
	"github.com/oggy107/rss-aggregator/utils"
)

func Start(db *database.Queries, concurrency int, timeBetweenRequests time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			utils.LogNonFatal(fmt.Sprintf("[Scraper]: %v", err))
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapFeed(db, wg, feed)
		}

		wg.Wait()
	}
}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		utils.LogNonFatal(fmt.Sprintf("[Scraper][failed to mark fetch]: %v", err))
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		utils.LogNonFatal(fmt.Sprintf("[Scraper][failed to fetch]: %v", err))
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			utils.LogNonFatal(fmt.Sprintf("[Scraper][date parser]: %v", err))
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Description: description,
			PublishedAt: pubDate,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}

			utils.LogNonFatal(fmt.Sprintf("[Scraper][failed to create post]: %v", err))
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
