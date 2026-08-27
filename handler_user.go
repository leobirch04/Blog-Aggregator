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
