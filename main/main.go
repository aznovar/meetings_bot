package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/db"
	"schedulertgbot/handlers"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_API_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Run database migrations
	db.RunMigrations()

	database, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "add_meeting":
			handlers.HandleAddMeeting(bot, update.Message, database)
		case "view_meetings":
			handlers.HandleViewMeetings(bot, update.Message, database)
		case "add_summary":
			handlers.HandleAddSummary(bot, update.Message, database)
		case "remind":
			handlers.HandleRemind(bot, update.Message, database)
		}
	}
}
