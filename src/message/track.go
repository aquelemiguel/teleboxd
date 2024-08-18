package message

import (
	"teleboxd/src/feed"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendTrackSuccess(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := BuildTrackSuccess(handle)
	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: false,
		},
	})
}

func SendNewFilmMessage(b *gotgbot.Bot, chatId int64, diary feed.LBDiary, item feed.LBItem) (*gotgbot.Message, error) {
	var message string

	if !item.Rewatch {
		if item.MemberRating != 0 {
			message = BuildNewFilmWatchRating(diary.MemberHandle, item)
		} else {
			message = BuildNewFilmWatch(diary.MemberHandle, item)
		}
	} else {
		if item.MemberRating != 0 {
			message = BuildNewFilmRewatchRating(diary.MemberHandle, item)
		} else {
			message = BuildNewFilmRewatch(diary.MemberHandle, item)
		}
	}

	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			PreferSmallMedia: true,
			Url:              item.FilmUrl,
		},
	})
}

func SendAlreadyTracking(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := BuildTrackDuplicateUser(handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidUser(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := BuildTrackInvalidUser(handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidTrackUsage(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	message := BuildTrackBadUsage()
	return SendMessage(b, chatId, message, nil)
}
