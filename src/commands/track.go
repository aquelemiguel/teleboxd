package commands

import (
	"fmt"
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
	}

	_, err := database.CreateMember(args[1], ctx.EffectiveChat.Id)
	if err != nil {
		message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.AlreadyTracking, args[1]))
		return nil
	}

	message.SendMessage(b, ctx.EffectiveChat.Id, fmt.Sprintf(locales.TrackSuccess, args[1]))
	return nil
}
