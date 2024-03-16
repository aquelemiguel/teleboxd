package message

import (
	"fmt"
	"groundhog/src/feed"
	"groundhog/src/locales"
	"math"
	s "strings"
)

func BuildNewFilmEntryMessage(diary feed.LBDiary, item feed.LBItem) string {
	fullStars := int(math.Floor(item.MemberRating))
	halfStars := int(math.Round(item.MemberRating - float64(fullStars)))
	stars := s.Repeat("★", fullStars) + s.Repeat("½", halfStars)

	s := fmt.Sprintf(
		locales.NewFilmEntry,
		diary.MemberLink,
		diary.MemberName,
		item.FilmUrl,
		item.FilmTitle,
		item.FilmYear,
		stars,
	)
	return s
}
