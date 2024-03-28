package message

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendMessage(b *gotgbot.Bot, chatId int64, message string, opts *gotgbot.SendMessageOpts) (*gotgbot.Message, error) {
	if opts == nil {
		return send(b, chatId, message, &gotgbot.SendMessageOpts{
			ParseMode: "HTML",
			LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
				IsDisabled: true,
			},
		})
	}
	return send(b, chatId, message, opts)
}

func send(b *gotgbot.Bot, chatId int64, message string, opts *gotgbot.SendMessageOpts) (*gotgbot.Message, error) {
	m, err := b.SendMessage(chatId, message, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return m, nil
}
