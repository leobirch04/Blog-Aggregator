package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func addFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("2 arguments expected")
	}

	userName := s.cfg.CurrentUserName
	userList, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}
	var currentUserId uuid.UUID

	for _, user := range userList {
		if user.Name == userName {
			currentUserId = user.ID
		}
	}

	params := database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], UserID: currentUserId}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("%v : %v : %v : %v", feed.ID, feed.UserID, feed.CreatedAt, feed.Name)

	return nil
}

func getAllFeeds(s *state, cmd command) error
{

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%v : %v : %v : %v", feed.ID, feed.UserID, feed.CreatedAt, feed.Name)
	}
	return nil
}
