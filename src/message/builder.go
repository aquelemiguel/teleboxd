package message

import (
	"fmt"
	"math"
	s "strings"
	"teleboxd/src/feed"
	"teleboxd/src/locales"
)

func BuildTrackSuccess(handle string) string {
	template, _ := locales.Translate("track_success")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🎥",
		"handle": buildHandleLink(handle),
	})
}

func BuildTrackInvalidUser(handle string) string {
	template, _ := locales.Translate("track_invalid_user")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
		"handle": buildHandleLink(handle),
	})
}

func BuildTrackDuplicateUser(handle string) string {
	template, _ := locales.Translate("track_duplicate_user")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
		"handle": buildHandleLink(handle),
	})
}

func BuildTrackBadUsage() string {
	template, _ := locales.Translate("track_bad_usage")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
		"usage": "<code>/track &lt;handle&gt;</code>",
	})
}

func BuildUntrackSuccess(handle string) string {
	template, _ := locales.Translate("untrack_success")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🗑️",
		"handle": buildHandleLink(handle),
	})
}

func BuildUntrackNotTrackingUser(handle string) string {
	template, _ := locales.Translate("untrack_not_tracking_user")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
		"handle": buildHandleLink(handle),
	})
}

func BuildUntrackBadUsage() string {
	template, _ := locales.Translate("untrack_bad_usage")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
		"usage": "<code>/untrack &lt;handle&gt;</code>",
	})
}

func BuildNewFilmWatch(handle string, film feed.LBItem) string {
	template, _ := locales.Translate("new_film_watch")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🍿",
		"handle": buildHandleLink(handle),
		"film": buildFilmLink(film.FilmUrl, film.FilmTitle, film.FilmYear),
	})
}

func BuildNewFilmWatchRating(handle string, film feed.LBItem) string {
	template, _ := locales.Translate("new_film_watch_rating")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🍿",
		"handle": buildHandleLink(handle),
		"film": buildFilmLink(film.FilmUrl, film.FilmTitle, film.FilmYear),
		"rating": buildRatingStars(film.MemberRating),
	})
}

func BuildNewFilmRewatch(handle string, film feed.LBItem) string {
	template, _ := locales.Translate("new_film_rewatch")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🔁",
		"handle": buildHandleLink(handle),
		"film": buildFilmLink(film.FilmUrl, film.FilmTitle, film.FilmYear),
	})
}

func BuildNewFilmRewatchRating(handle string, film feed.LBItem) string {
	template, _ := locales.Translate("new_film_rewatch_rating")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🔁",
		"handle": buildHandleLink(handle),
		"film": buildFilmLink(film.FilmUrl, film.FilmTitle, film.FilmYear),
		"rating": buildRatingStars(film.MemberRating),
	})
}

func BuildListHeader() string {
	template, _ := locales.Translate("list_header")

	return replacePlaceholders(template, map[string]string{
		"emoji": "&#128466;&#65039;",
	})
}

func BuildListHeaderEmpty() string {
	template, _ := locales.Translate("list_header_empty")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🎥",
	})
}

func BuildListEntry(handle string) string {
	template, _ := locales.Translate("list_entry")

	return replacePlaceholders(template, map[string]string{
		"handle": buildHandleLink(handle),
	})
}

func BuildSomethingWentWrong() string {
	template, _ := locales.Translate("something_went_wrong")

	return replacePlaceholders(template, map[string]string{
		"emoji": "🚫",
	})
}

func replacePlaceholders(message string, args map[string]string) string {
	for key, value := range args {
		message = s.ReplaceAll(message, fmt.Sprintf("{%s}", key), value)
	}
	return message
}

func buildHandleLink(handle string) string {
	return fmt.Sprintf("<b><a href=\"https://letterboxd.com/%s\">%s</a></b>", handle, handle)
}

func buildFilmLink(url string, title string, year string) string {
	return fmt.Sprintf("<b><a href=\"%s\">%s (%s)</a></b>", url, title, year)
}

func buildRatingStars(rating float64) string {
	full := int(math.Floor(rating))
	half := int(math.Round(rating - float64(full)))
	return s.Repeat("★", full) + s.Repeat("½", half)
}