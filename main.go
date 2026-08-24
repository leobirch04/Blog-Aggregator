package main

import (
	"blog-aggregator/internal/config"
	"log"
	"os"
)

type state struct {
	config *config.JsonConfig
}

func main() {
	con, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	sta := state{&con}
	cmds := commands{
		CommandMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	err = cmds.run(&sta, command{args[1], args[2:]})
	if err != nil {
		log.Fatal(err)
	}
}
