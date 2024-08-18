package message

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendSomethingWentWrong(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	message := BuildSomethingWentWrong()
	return SendMessage(b, chatId, message, nil)
}
