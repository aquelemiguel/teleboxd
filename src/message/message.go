package message

import (
	"fmt"
	"groundhog/src/database"
	"groundhog/src/feed"
	"groundhog/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func SendNewFilmMessage(b *gotgbot.Bot, chatId int64, diary feed.LBDiary, item feed.LBItem) (*gotgbot.Message, error) {
	message := BuildNewFilmEntryMessage(diary, item)

	return send(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			PreferSmallMedia: true,
			Url:              item.FilmUrl,
		},
	})
}

func SendListMessage(b *gotgbot.Bot, ctx *ext.Context, users []*database.User) (*gotgbot.Message, error) {
	chatId := ctx.EffectiveChat.Id

	// build the message
	message := locales.ListHeader
	if len(users) == 0 {
		message = locales.ListHeaderEmpty
	}
	for _, user := range users {
		message += fmt.Sprintf(locales.ListEntry, user.Handle, user.Handle)
	}

	return send(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: true,
		},
	})
}

func SendMessage(b *gotgbot.Bot, chatId int64, message string) (*gotgbot.Message, error) {
	return send(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
	})
}

func send(b *gotgbot.Bot, chatId int64, message string, opts *gotgbot.SendMessageOpts) (*gotgbot.Message, error) {
	m, err := b.SendMessage(chatId, message, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return m, nil
}
