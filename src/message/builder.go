package message

import (
	"fmt"
	"groundhog/src/feed"
	"groundhog/src/locales"
	"math"
	s "strings"
)

func BuildNewTrackEntryMessage(item feed.LetterboxdItem) string {
	fullStars := int(math.Floor(item.MemberRating))
	halfStars := int(math.Round(item.MemberRating - float64(fullStars)))
	stars := s.Repeat("★", fullStars) + s.Repeat("½", halfStars)

	s := fmt.Sprintf(
		locales.NewTrackEntry,
		"aquelemiguel",
		item.FilmTitle,
		item.FilmUrl,
		item.FilmYear,
		stars,
	)
	return s
}
