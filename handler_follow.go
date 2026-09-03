package main

import (
	"blog-aggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func addFollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 {
		return errors.New("1 arguments expected")
	}

	feedID, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID.ID,
	}

	feed, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("feed: %v\nuser: %v\n", feed.FeedName, feed.UserName)

	return nil
}

func getFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("no arguments expected")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return err
	}
	fmt.Println("all feeds being followed: ")
	for _, follow := range follows {
		fmt.Printf("%v\n", follow.FeedName)
	}
	return nil
}

func removeFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("one argument expected")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	params := database.DeleteFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println("unfollowed")
	return nil
}
