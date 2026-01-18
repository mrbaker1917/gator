package main

import (
	"fmt"

	"github.com/mrbaker1917/gator/internal/config"
	"github.com/mrbaker1917/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command %s not found", cmd.name)
	}

	return f(s, cmd)
}
