package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"teleboxd/src/commands"
	"teleboxd/src/core"
	"teleboxd/src/database"
	"teleboxd/src/locales"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN env var must be set")
	}

	// initialize the database
	_, err := database.GetDatabase()
	if err != nil {
		log.Fatal("failed to initialize database:", err.Error())
	}

	// initialize the locales
	// todo: fetch the correct locale based on the chat setting
	err = locales.LoadLocales()
	if err != nil {
		log.Fatal("failed to initialize locale:", err.Error())
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

	// simple health check to help with monitoring
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Printf("%s has been started", b.User.Username)

	// revive tracking for users already in the db
	core.Revive(b)

	updater.Idle()
}
