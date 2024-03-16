package commands

import (
	"errors"
	"fmt"
	"groundhog/src/core"
	"groundhog/src/database"
	"groundhog/src/locales"
	"groundhog/src/message"
	s "strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Untrack(b *gotgbot.Bot, ctx *ext.Context) error {
	args := s.Split(ctx.EffectiveMessage.Text, " ")

	if len(args) != 2 {
		message.SendMessage(b, ctx.EffectiveChat.Id, locales.InvalidUntrackUsage)
		return nil
	}
	handle := args[1]

	err := database.DeleteMember(handle, ctx.EffectiveChat.Id)
	if errors.Is(err, database.ErrUserNotFound) {
		message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.NotTracking, handle))
		return nil
	}

	// if this is the user's last chat, stop polling them
	chats, _ := database.GetChatsByUser(handle)
	if len(chats) == 0 {
		core.StopPolling(handle)
	}

	message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.UntrackSuccess, handle))
	return nil
}
