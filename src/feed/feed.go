package feed

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/mmcdole/gofeed"
)

type LBDiary struct {
	MemberName   string
	MemberHandle string
	MemberLink   string
	Items        []*LBItem
}

type LBItem struct {
	FilmTitle    string
	FilmUrl      string
	FilmYear     string
	MemberRating float64
	Rewatch      bool
	WatchedAt    int64
}

func Fetch(handle string) *LBDiary {
	url := fmt.Sprintf("https://letterboxd.com/%s/rss/", handle)
	fp := gofeed.NewParser()

	f, err := fp.ParseURL(url)
	if err != nil {
		fmt.Println("failed to fetch feed:", err.Error())
		return &LBDiary{}
	}

	var items []*LBItem
	for _, item := range f.Items {
		if !s.HasPrefix(item.GUID, "letterboxd-watch") {
			continue
		}

		ext := item.Extensions["letterboxd"]
		lbi := &LBItem{
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
	}
	diary := &LBDiary{
		MemberName:   s.Split(f.Title, " - ")[1],
		MemberHandle: s.Split(f.Link, "/")[3],
		MemberLink:   f.Link,
		Items:        items,
	}
	// j, _ := json.MarshalIndent(diary, "", "  ")
	// fmt.Println(string(j))
	return diary
}
