package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/db"
)

func HandleRemind(bot *tgbotapi.BotAPI, message *tgbotapi.Message, repo *db.MeetingRepository) {
	meetings, err := repo.GetUpcomingMeetings()
	if err != nil {
		log.Fatal(err)
	}

	var response string
	for _, meeting := range meetings {
		response += fmt.Sprintf("Напоминание! Встреча \"%s\" с участниками %s состоится %s\n\n",
			meeting.Title, meeting.Participants, meeting.Date.Format("2006-01-02 15:04"))
	}

	if response == "" {
		response = "Нет встреч в ближайшие 24 часа."
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
