package message

import (
	"fmt"
	"teleboxd/src/database"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

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
