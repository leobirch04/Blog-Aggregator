package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func addFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("2 arguments expected")
	}

	params := database.CreateFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0], UserID: user.ID, Url: cmd.args[1]}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("%v : %v : %v : %v", feed.ID, feed.UserID, feed.CreatedAt, feed.Name)

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	return nil
}

func getAllFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		feedUsername, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("\n[%v : %v : %v : %v]\n", feed.ID, feedUsername, feed.Url, feed.Name)
	}
	return nil
}
