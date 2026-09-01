package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("one argument expected")
	}
	exists, err := s.db.UserExists(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	if !exists {
		os.Exit(1)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Username Set...")

	return nil
}

func handlerGetAllUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments expected")
	}

	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		current := ""
		if s.cfg.CurrentUserName == user.Name {
			current = "(current)"
		}
		fmt.Printf("* %v %v", user.Name, current)
	}
	return nil
}
