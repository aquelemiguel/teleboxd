package message

import (
	"fmt"
	"teleboxd/src/database"
	"teleboxd/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

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

	return SendMessage(b, chatId, message, &gotgbot.SendMessageOpts{
		ParseMode: "HTML",
		LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
			IsDisabled: true,
		},
	})
}
