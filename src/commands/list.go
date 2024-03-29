package commands

import (
	"teleboxd/src/database"
	"teleboxd/src/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func List(b *gotgbot.Bot, ctx *ext.Context) error {
	users, err := database.GetUsersByChat(ctx.EffectiveChat.Id)
	if err != nil {
		// TODO: properly handle this error
		return err
	}
	message.SendListMessage(b, ctx, users)
	return nil
}
