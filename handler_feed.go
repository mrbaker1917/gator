package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrbaker1917/gator/internal/database"
)

func createFeedFollowForUser(ctx context.Context, db *database.Queries, userID, feedID uuid.UUID) error {
	_, err := db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		UserID: userID,
		FeedID: feedID,
	})
	return err
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <feedname> <url>")
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	ctx := context.Background()

	feed, err := s.db.CreateFeed(
		ctx,
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      feedName,
			Url:       feedURL,
			UserID:    user.ID,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	if err := createFeedFollowForUser(ctx, s.db, user.ID, feed.ID); err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Println("ID: ", feed.ID)
	fmt.Println("Name: ", feed.Name)
	fmt.Println("URL: ", feed.Url)
	fmt.Println("UserID: ", feed.UserID)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("we encountered an error: %w", err)
	}
	fmt.Println("Here are all the feeds:")
	for _, feed := range feeds {
		fmt.Println("-------------------------------------------")
		fmt.Printf("Feed Name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed Creator: %s\n", feed.Name_2)
		fmt.Println("-------------------------------------------")
	}
	return nil
}
