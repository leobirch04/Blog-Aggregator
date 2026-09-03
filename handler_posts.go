package main

import (
	"blog-aggregator/internal/database"
	"errors"
	"fmt"
	"strconv"
)

func browse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return errors.New("one argument expected")
	}

	limit := 2

	if len(cmd.args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}

		limit = parsedLimit
	}

	for i := 0; i < limit; i++ {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

	return nil
}
