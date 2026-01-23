package main

import (
	"context"
	"fmt"

	"github.com/mrbaker1917/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed for this url: %w", err)
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

func handlerFollowing(s *state, cmd command, user database.User) error {

	ctx := context.Background()
	rows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Could not get feeds: %w", err)
	}
	for _, row := range rows {
		fmt.Println(row.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	ctx := context.Background()
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed for this url: %w", err)
	}

	if err = s.db.UnfollowFeed(ctx, database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}); err != nil {
		return fmt.Errorf("could not unfollow from this feed: %w", err)
	}

	return nil
}
