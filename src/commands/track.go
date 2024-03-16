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

func Track(b *gotgbot.Bot, ctx *ext.Context) error {
	args := s.Split(ctx.EffectiveMessage.Text, " ")

	if len(args) != 2 {
		message.SendMessage(b, ctx.EffectiveChat.Id, locales.InvalidTrackUsage)
		return nil
	}

	_, err := database.CreateMember(args[1], ctx.EffectiveChat.Id)
	if errors.Is(err, database.ErrUserNotFound) {
		message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.AlreadyTracking, args[1]))
		return nil
	}

	// if this is a fresh user, start polling them
	ticker := core.GetUserTicker(args[1])
	if ticker == nil {
		core.StartPolling(b, args[1])
	}

	message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.TrackSuccess, args[1]))
	return nil
}
