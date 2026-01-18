package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mrbaker1917/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: register <username>")
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
		},
	)

	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}
	fmt.Println("User created successfully!")
	fmt.Printf("ID: %v\n", user.ID)
	fmt.Printf("Name: %v\n", user.Name)
	return nil

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: login <username>")
	}
	username := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		return fmt.Errorf("could not login user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	fmt.Printf("User set to %s\n", user.Name)
	return nil
}

func reset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteAll(ctx)
	if err != nil {
		return fmt.Errorf("we encountered an error: %w", err)
	}
	fmt.Println("All users were deleted.")
	return nil
}

func users(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("we encountered an error: %w", err)
	}
	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
