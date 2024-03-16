package main

import (
	"log"
	"os"
	"time"

	"groundhog/src/commands"
	"groundhog/src/core"
	"groundhog/src/database"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN env var must be set")
	}

	// initialize the database
	_, err = database.GetDatabase()
	if err != nil {
		log.Fatal("failed to initialize database:", err.Error())
	}

	// initialize the telegram bot
	b, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("failed to create a new bot", err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	updater := ext.NewUpdater(dispatcher, nil)

	dispatcher.AddHandler(handlers.NewCommand("list", commands.List))
	dispatcher.AddHandler(handlers.NewCommand("track", commands.Track))
	dispatcher.AddHandler(handlers.NewCommand("untrack", commands.Untrack))

	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		log.Fatal("failed to start polling:", err.Error())
	}
	log.Printf("%s has been started", b.User.Username)

	// revive tracking for users already in the db
	core.Revive(b)

	updater.Idle()
}
