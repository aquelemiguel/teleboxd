package feed

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/mmcdole/gofeed"
)

// TODO: make a struct for []*LetterboxdItem
type LetterboxdItem struct {
	FilmTitle    string
	FilmUrl      string
	FilmYear     string
	MemberRating float64
	Rewatch      bool
	WatchedAt    int64
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
			MemberRating: func() float64 {
				rating, _ := strconv.ParseFloat(ext["memberRating"][0].Value, 64)
				return rating
			}(),
			Rewatch:   ext["rewatch"][0].Value == "Yes",
			WatchedAt: item.PublishedParsed.Unix(),
		}
		items = append(items, lbi)

		// j, _ := json.MarshalIndent(lbi, "", "  ")
		// fmt.Println(string(j))
	}
	return items
}
