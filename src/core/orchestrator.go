package core

import (
	"groundhog/src/database"
	"groundhog/src/feed"
	"groundhog/src/message"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func Revive(b *gotgbot.Bot) {
	// fetch all users in the database
	users, err := database.GetAllUsers()
	if err != nil {
		// TODO: handle this error
		return
	}

	// start polling for each user
	for _, user := range users {
		log.Printf("reviving polling for user @%s...\n", user.Handle)
		StartPolling(b, user.Handle)
	}
}

func StartPolling(b *gotgbot.Bot, handle string) *time.Ticker {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for now := range ticker.C {
			// TODO: handle this error
			f := feed.Fetch(handle)

			// fetch the last polling time
			_, err := database.GetUser(handle)
			if err != nil {
				// TODO: implement retries in the future
				continue
			}

			// use it to filter the items by unseen
			var unseen []*feed.LetterboxdItem
			for _, item := range f {
				if item.WatchedAt > now.AddDate(0, 0, -7).Unix() {
					unseen = append(unseen, item)
				}
			}

			// if there are no new items, we're done here
			if len(unseen) == 0 {
				log.Printf("no new items for user @%s", handle)
				continue
			}

			// fetch chats that are subscribed to this feed
			chats, err := database.GetChatsByUser(handle)
			if err != nil {
				// TODO: implement retries in the future
				continue
			}

			// send the new items to the chats
			for _, chatId := range chats {
				for _, item := range unseen {
					// TODO: handle fail states here
					message.SendNewFilmMessage(b, chatId, *item)
				}
			}

			// TODO: this should only be updated if all messages were sent successfully
			database.UpdateUser(handle, now.Unix())
		}
	}()
	return ticker
}
