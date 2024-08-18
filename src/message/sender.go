package message

import (
	"fmt"
	"teleboxd/src/database"
	"teleboxd/src/feed"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
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

func SendUntrackSuccess(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := BuildUntrackSuccess(handle)
	return SendMessage(b, chatId, message, nil)
}

func SendNotTracking(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := BuildUntrackNotTrackingUser(handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidUntrackUsage(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	message := BuildUntrackBadUsage()
	return SendMessage(b, chatId, message, nil)
}

func SendListMessage(b *gotgbot.Bot, ctx *ext.Context, users []*database.User) (*gotgbot.Message, error) {
	var message string
	chatId := ctx.EffectiveChat.Id

	if len(users) > 0 {
		message = fmt.Sprintf("%s\n", BuildListHeader())
	} else {
		message = fmt.Sprintf("%s\n", BuildListHeaderEmpty())
	}
	for _, user := range users {
		message += fmt.Sprintf("%s\n", BuildListEntry(user.Handle))
	}
	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: true,
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

func SendSomethingWentWrong(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	message := BuildSomethingWentWrong()
	return SendMessage(b, chatId, message, nil)
}
