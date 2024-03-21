package core

import (
	"groundhog/src/database"
	"groundhog/src/feed"
	"groundhog/src/message"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

var pool = make(map[string]*time.Ticker)

func StartPolling(b *gotgbot.Bot, handle string) *time.Ticker {
	ticker := time.NewTicker(30 * time.Minute)
	pool[handle] = ticker

	go func() {
		for now := range ticker.C {
			// TODO: handle this error
			f, _ := feed.Fetch(handle)
			log.Printf("fetched %d items for user @%s", len(f.Items), handle)

			// fetch the last polling time
			_, err := database.GetUser(handle)
			if err != nil {
				// TODO: implement retries in the future
				continue
			}
			// use it to filter the items by unseen
			var unseen []*feed.LBItem
			for _, item := range f.Items {
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
					message.SendNewFilmMessage(b, chatId, *f, *item)
				}
			}
			// TODO: this should only be updated if all messages were sent successfully
			database.UpdateUser(handle, now.Unix())
		}
	}()
	return ticker
}

func StopPolling(handle string) bool {
	ticker := pool[handle]
	if ticker == nil {
		return false
	}
	ticker.Stop()
	delete(pool, handle)
	return true
}

func GetUserTicker(handle string) *time.Ticker {
	return pool[handle]
}
