package message

import (
	"fmt"
	"groundhog/src/feed"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendNewFilmMessage(b *gotgbot.Bot, chatId int64, item feed.LetterboxdItem) (*gotgbot.Message, error) {
	message := BuildNewFilmEntryMessage(item)

	return send(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			PreferSmallMedia: true,
			Url:              item.FilmUrl,
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
