package rss

import (
	"encoding/json"
	"fmt"
	"strconv"
	s "strings"

	"github.com/mmcdole/gofeed"
)

type ExtendedItem struct {
	*gofeed.Item
	FilmTitle    string
	FilmYear     string
	MemberRating float32
	Rewatch      bool
}

type ExtendedFeed struct {
	Feed  *gofeed.Feed
	Items []*ExtendedItem
}

func FetchFeed(handle string) {
	url := fmt.Sprintf("https://letterboxd.com/%s/rss/", handle)
	fp := gofeed.NewParser()

	f, err := fp.ParseURL(url)
	if err != nil {
		fmt.Println("failed to fetch feed:", err.Error())
		return
	}

	ef, _ := ExtendFeed(f)

	for _, movie := range ef.Items {
		j, _ := json.MarshalIndent(movie, "", "  ")
		fmt.Println(string(j))
	}
}

func ExtendFeed(f *gofeed.Feed) (*ExtendedFeed, error) {
	ef := &ExtendedFeed{
		Feed:  f,
		Items: make([]*ExtendedItem, len(f.Items)),
	}

	var items []*ExtendedItem
	for _, item := range f.Items {
		if !s.HasPrefix(item.GUID, "letterboxd-watch") {
			continue
		}

		ext := item.Extensions["letterboxd"]

		ef := &ExtendedItem{
			Item:      item,
			FilmTitle: ext["filmTitle"][0].Value,
			FilmYear:  ext["filmYear"][0].Value,
			MemberRating: func() float32 {
				rating, _ := strconv.ParseFloat(ext["memberRating"][0].Value, 32)
				return float32(rating)
			}(),
			Rewatch: ext["rewatch"][0].Value == "Yes",
		}

		items = append(items, ef)
	}
	ef.Items = items

	return ef, nil
}
