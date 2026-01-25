package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mrbaker1917/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = n
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			ID:    user.ID,
			Limit: int32(limit),
		},
	)
	if err != nil {
		return fmt.Errorf("could not get posts for user %s: %w", user.Name, err)
	}

	for _, post := range posts {
		fmt.Printf("%s - %s\n",
			post.PublishedAt.Format("Mon Jan 2"),
			post.Name,
		)

		fmt.Printf("TITLE: --- %s ---\n", post.Title)

		if post.Description.Valid {
			fmt.Printf("Description:    %s\n", post.Description.String)
		}

		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("====================================")
	}

	return nil
}
