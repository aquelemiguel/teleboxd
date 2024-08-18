package message

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

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
