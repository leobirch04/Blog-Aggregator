package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
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
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
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
		fmt.Printf("\ntitle: %v\ndescription: %v", item.Title, item.Description)
	}
	return nil
}
