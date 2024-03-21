package commands

import (
	"errors"
	"fmt"
	"groundhog/src/core"
	"groundhog/src/database"
	"groundhog/src/feed"
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
	handle := args[1]

	// ensure the user is a valid Letterboxd user
	_, err := feed.Fetch(handle)
	if err != nil {
		message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.TrackInvalidUser, handle, handle))
		return nil
	}

	_, err = database.CreateMember(handle, ctx.EffectiveChat.Id)
	if errors.Is(err, database.ErrUserAlreadyExists) {
		message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.AlreadyTracking, handle, handle))
		return nil
	}

	// if this is a fresh user, start polling them
	ticker := core.GetUserTicker(handle)
	if ticker == nil {
		core.StartPolling(b, handle)
	}

	message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.TrackSuccess, handle, handle))
	return nil
}
