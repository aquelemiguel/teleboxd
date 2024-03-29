package message

import (
	"fmt"
	"teleboxd/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func SendUntrackSuccess(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := fmt.Sprintf(locales.UntrackSuccess, handle, handle)
	return SendMessage(b, chatId, message, nil)
}

func SendNotTracking(b *gotgbot.Bot, chatId int64, handle string) (*gotgbot.Message, error) {
	message := fmt.Sprintf(locales.NotTracking, handle, handle)
	return SendMessage(b, chatId, message, nil)
}

func SendInvalidUntrackUsage(b *gotgbot.Bot, chatId int64) (*gotgbot.Message, error) {
	return SendMessage(b, chatId, locales.InvalidUntrackUsage, nil)
}
