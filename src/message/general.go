package message

import (
	"groundhog/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendSomethingWentWrong(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	return SendMessage(b, chatId, locales.SomethingWentWrong, nil)
}
