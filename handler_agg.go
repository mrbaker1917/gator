package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

}

func scrapeFeeds(s *state) {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Println("could not get next feed: ", err)
		return
	}
	fmt.Printf("Fetching from %s RSS\n", feed.Name)
	scrapeFeed(ctx, s.db, feed)
}
