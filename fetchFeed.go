package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", `gator"`)
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	result := RSSFeed{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return &RSSFeed{}, err
	}

	return &result, nil
}

func printFeed(s *state, cmd command) error {
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func scrapeFeeds(s *state) error {

	nextfeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), nextfeed.ID)
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), nextfeed.Url)
	if err != nil {
		return err
	}

	title := html.UnescapeString(feed.Channel.Title)
	des := html.UnescapeString(feed.Channel.Description)
	fmt.Printf("title: %v\ndescription: %v", title, des)

	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Link = html.UnescapeString(item.Link)
		item.Description = html.UnescapeString(item.Description)
		item.PubDate = html.UnescapeString(item.PubDate)
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: time.Now(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			FeedID:      nextfeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			return err
		}

	}
	return nil
}
