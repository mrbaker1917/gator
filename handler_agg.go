package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mrbaker1917/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeed(ctx context.Context, db *database.Queries, feed database.Feed) {
	err := db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		log.Printf("could not mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("could not fetch feed %s: %v", feed.Name, err)
		return
	}

	const rssTimeLayout = "Mon, 02 Jan 2006 15:04:05 -0700"

	for _, item := range feedData.Channel.Item {
		t, err := time.Parse(rssTimeLayout, item.PubDate)
		if err != nil {
			t = time.Now().UTC()
		}
		desc := sql.NullString{}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		_, err = db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: t,
			FeedID:      feed.ID,
		},
		)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			log.Printf("could not post feed item: %v", err)
		}
	}

}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Println("could not get next feed: ", err)
		return
	}
	fmt.Printf("----- Fetching stories from %s RSS -----\n", feed.Name)
	scrapeFeed(ctx, s.db, feed)
}
