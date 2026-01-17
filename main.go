package main

import (
	"fmt"
	"os"

	"github.com/mrbaker1917/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	st := state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	name := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		name: name,
		args: args,
	}

	if err := cmds.run(&st, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
