package message

import (
	"fmt"
	"math"
	s "strings"
	"teleboxd/src/feed"
	"teleboxd/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendTrackSuccess(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := fmt.Sprintf(locales.TrackSuccess, handle, handle)
	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: false,
		},
	})
}

func SendNewFilmMessage(b *gotgbot.Bot, chatId int64, diary feed.LBDiary, item feed.LBItem) (*gotgbot.Message, error) {
	var message, template string
	var stars string

	if item.MemberRating != 0 {
		full := int(math.Floor(item.MemberRating))
		half := int(math.Round(item.MemberRating - float64(full)))
		stars = s.Repeat("★", full) + s.Repeat("½", half)
	}
	if item.Rewatch {
		if stars != "" {
			template = locales.NewFilmRewatchRating
		} else {
			template = locales.NewFilmRewatch
		}
	} else {
		if stars != "" {
			template = locales.NewFilmWatchRating
		} else {
			template = locales.NewFilmWatch
		}
	}

	args := []interface{}{
		diary.MemberLink, diary.MemberName, item.FilmUrl, item.FilmTitle, item.FilmYear,
	}
	if stars != "" {
		args = append(args, stars)
	}
	message = fmt.Sprintf(template, args...)

	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			PreferSmallMedia: true,
			Url:              item.FilmUrl,
		},
	})
}

func SendAlreadyTracking(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := fmt.Sprintf(locales.AlreadyTracking, handle, handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidUser(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := fmt.Sprintf(locales.TrackInvalidUser, handle, handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidTrackUsage(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	return SendMessage(b, chatId, locales.InvalidTrackUsage, nil)
}
