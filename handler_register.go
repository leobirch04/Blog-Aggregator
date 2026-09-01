package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("one argument expected")
	}

	exists, err := s.db.UserExists(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("d")
	if exists {
		os.Exit(1)
	}

	params := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0]}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user created")
	log.Printf("[%v: %v: %v: %v]", user.ID, user.Name, user.CreatedAt, user.UpdatedAt)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All Users Cleared...")

	err = s.db.ClearFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All Feeds Cleared...")

	return nil
}
