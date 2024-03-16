package message

import (
	"fmt"
	"groundhog/src/feed"
	"groundhog/src/locales"
	"math"
	s "strings"
)

func BuildNewFilmEntryMessage(item feed.LBItem) string {
	fullStars := int(math.Floor(item.MemberRating))
	halfStars := int(math.Round(item.MemberRating - float64(fullStars)))
	stars := s.Repeat("★", fullStars) + s.Repeat("½", halfStars)

	s := fmt.Sprintf(
		locales.NewFilmEntry,
		// TODO: replace this with the user's handle
		"https://letterboxd.com/aquelemiguel/",
		"aquelemiguel",
		item.FilmUrl,
		item.FilmTitle,
		item.FilmYear,
		stars,
	)
	return s
}
