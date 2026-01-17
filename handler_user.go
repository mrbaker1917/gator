package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: login <username>")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	fmt.Printf("User set to %s\n", username)

	return nil
}
