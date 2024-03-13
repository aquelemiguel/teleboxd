package feed

import (
	"encoding/json"
	"fmt"
	"groundhog/src/database"
	"strconv"
	s "strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type LetterboxdItem struct {
	FilmTitle    string
	FilmUrl      string
	FilmYear     string
	MemberRating float32
	Rewatch      bool
	WatchedAt    int64
}

func StartPoll(handle string) *time.Ticker {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for now := range ticker.C {
			fmt.Println("fetching feed at", now)
			database.UpdateUser(handle, now.Unix())
		}
	}()
	return ticker
}

func Fetch(handle string) []*LetterboxdItem {
	url := fmt.Sprintf("https://letterboxd.com/%s/rss/", handle)
	fp := gofeed.NewParser()

	f, err := fp.ParseURL(url)
	if err != nil {
		fmt.Println("failed to fetch feed:", err.Error())
		return []*LetterboxdItem{}
	}

	var items []*LetterboxdItem

	for _, item := range f.Items {
		if !s.HasPrefix(item.GUID, "letterboxd-watch") {
			continue
		}

		ext := item.Extensions["letterboxd"]
		lbi := &LetterboxdItem{
			FilmTitle: ext["filmTitle"][0].Value,
			FilmUrl: func() string {
				id := s.Split(item.Link, "/")[5]
				return fmt.Sprintf("https://letterboxd.com/film/%s", id)
			}(),
			FilmYear: ext["filmYear"][0].Value,
			MemberRating: func() float32 {
				rating, _ := strconv.ParseFloat(ext["memberRating"][0].Value, 32)
				return float32(rating)
			}(),
			Rewatch:   ext["rewatch"][0].Value == "Yes",
			WatchedAt: item.PublishedParsed.Unix(),
		}
		items = append(items, lbi)

		j, _ := json.MarshalIndent(lbi, "", "  ")
		fmt.Println(string(j))
	}
	return items
}
