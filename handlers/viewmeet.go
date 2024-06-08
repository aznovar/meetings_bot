package handlers

import (
	"database/sql"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/models"
)

func HandleViewMeetings(bot *tgbotapi.BotAPI, message *tgbotapi.Message,
	db *sql.DB) {
	rows, err := db.Query("SELECT id, title, date, participants, summary FROM meetings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var response string
	for rows.Next() {
		var meeting models.Meeting
		rows.Scan(&meeting.ID, &meeting.Title, &meeting.Date, &meeting.Participants, &meeting.Summary)
		response += fmt.Sprintf("ID: %d\nTitle: %s\nDate: %s\nParticipants: %s\nSummary: %s\n\n",
			meeting.ID, meeting.Title, meeting.Date.Format("2006-01-02 15:04"), meeting.Participants, meeting.Summary)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if response == "" {
		response = "Нет запланированных встреч."
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
