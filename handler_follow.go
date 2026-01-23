package main

import (
	"context"
	"fmt"

	"github.com/mrbaker1917/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	username := s.cfg.CurrentUserName
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("could not find user: %w", err)
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed for this feed: %w", err)
	}

	row, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not link user to feed: %w", err)
	}

	fmt.Println("You are now following this feed:")
	fmt.Printf("Feed: %s\n", row.FeedName)
	fmt.Printf("User: %s\n", row.UserName)

	return nil

}

func handlerFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUserName
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return fmt.Errorf("could not find user: %w", err)
	}
	rows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Could not get feeds: %w", err)
	}
	for _, row := range rows {
		fmt.Println(row.FeedName)
	}

	return nil
}
