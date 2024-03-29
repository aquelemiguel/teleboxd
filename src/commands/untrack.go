package commands

import (
	"errors"
	s "strings"
	"teleboxd/src/core"
	"teleboxd/src/database"
	"teleboxd/src/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Untrack(b *gotgbot.Bot, ctx *ext.Context) error {
	args := s.Split(ctx.EffectiveMessage.Text, " ")

	if len(args) != 2 {
		message.SendInvalidUntrackUsage(b, ctx.EffectiveChat.Id)
		return nil
	}
	handle := args[1]

	err := database.DeleteMember(handle, ctx.EffectiveChat.Id)
	if errors.Is(err, database.ErrUserNotFound) {
		message.SendNotTracking(b, ctx.EffectiveChat.Id, handle)
		return nil
	}

	// if this is the user's last chat, stop polling them
	chats, _ := database.GetChatsByUser(handle)
	if len(chats) == 0 {
		core.StopPolling(handle)
	}

	message.SendUntrackSuccess(b, ctx.EffectiveChat.Id, handle)
	return nil
}
