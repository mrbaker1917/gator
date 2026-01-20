package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrbaker1917/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	username := s.cfg.CurrentUserName
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <feedname> <url>")
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	ctx := context.Background()
	user, err := s.db.GetUser(ctx, username)

	if err != nil {
		return fmt.Errorf("could not login user: %w", err)
	}

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

	fmt.Println("ID: ", feed.ID)
	fmt.Println("Name: ", feed.Name)
	fmt.Println("URL: ", feed.Url)
	fmt.Println("UserID: ", feed.UserID)

	return nil
}
