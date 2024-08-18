package commands

import (
	"errors"
	s "strings"
	"teleboxd/src/core"
	"teleboxd/src/database"
	"teleboxd/src/feed"
	"teleboxd/src/message"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Track(b *gotgbot.Bot, ctx *ext.Context) error {
	args := s.Split(ctx.EffectiveMessage.Text, " ")

	if len(args) != 2 {
		message.SendInvalidTrackUsage(b, ctx.EffectiveChat.Id)
		return nil
	}
	handle := args[1]

	// ensure the user is a valid Letterboxd user
	_, err := feed.Fetch(handle)
	if err != nil {
		if err == feed.ErrUserDoesNotExist {
			message.SendInvalidUser(b, ctx.EffectiveChat.Id, handle)
			return nil
		}
		message.SendSomethingWentWrong(b, ctx.EffectiveChat.Id)
		return nil
	}

	_, err = database.CreateMember(handle, ctx.EffectiveChat.Id)
	if errors.Is(err, database.ErrUserAlreadyExists) {
		message.SendAlreadyTracking(b, ctx.EffectiveChat.Id, handle)
		return nil
	}

	// if this is a fresh user, start polling them
	ticker := core.GetUserTicker(handle)
	if ticker == nil {
		core.StartPolling(b, handle)
	}
	
	message.SendTrackSuccess(b, ctx.EffectiveChat.Id, handle)
	return nil
}
