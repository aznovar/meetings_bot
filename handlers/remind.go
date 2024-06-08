package handlers

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/models"
	"time"
)

func HandleRemind(bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *sql.DB) {
	rows, err := db.Query("SELECT id, title, date, participants FROM meetings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var response string
	currentTime := time.Now()
	for rows.Next() {
		var meeting models.Meeting
		rows.Scan(&meeting.ID, &meeting.Title, &meeting.Date, &meeting.Participants)
		if meeting.Date.Sub(currentTime).Hours() < 24 {
			response += fmt.Sprintf("Напоминание! Встреча \"%s\" с участниками %s состоится %s\n\n",
				meeting.Title, meeting.Participants, meeting.Date.Format("2006-01-02 15:04"))
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if response == "" {
		response = "Нет встреч в ближайшие 24 часа."
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
