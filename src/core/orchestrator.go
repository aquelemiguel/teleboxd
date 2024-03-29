package core

import (
	"log"
	"teleboxd/src/database"

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
