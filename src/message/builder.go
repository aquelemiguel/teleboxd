package message

import (
	"fmt"
	s "strings"
	"teleboxd/src/locales"
)

func BuildTrackSuccess(handle string) string {
	template, _ := locales.Translate("track_success")

	return replacePlaceholders(template, map[string]string{
		"emoji": "&#9989;",
		"handle": fmt.Sprintf("<b><a href=\"https://letterboxd.com/%s\">%s</a></b>", handle, handle),
	})
}

func BuildUntrackSuccess(handle string) string {
	template, _ := locales.Translate("untrack_success")

	return replacePlaceholders(template, map[string]string{
		"emoji": "&#9989;",
		"handle": fmt.Sprintf("<b><a href=\"https://letterboxd.com/%s\">%s</a></b>", handle, handle),
	})
}

func replacePlaceholders(message string, args map[string]string) string {
	for key, value := range args {
		message = s.ReplaceAll(message, fmt.Sprintf("{%s}", key), value)
	}
	return message
}