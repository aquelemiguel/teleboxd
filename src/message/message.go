package message

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendMessage(b *gotgbot.Bot, chatId int64, message string) (*gotgbot.Message, error) {
	m, err := b.SendMessage(chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "Markdown",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return m, nil
}
