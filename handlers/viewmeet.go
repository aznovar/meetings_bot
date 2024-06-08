package handlers

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/db"
)

func HandleViewMeetings(bot *tgbotapi.BotAPI, message *tgbotapi.Message, repo *db.MeetingRepository) {
	meetings, err := repo.GetMeetings()
	if err != nil {
		log.Fatal(err)
	}

	var response string
	for _, meeting := range meetings {
		response += fmt.Sprintf("ID: %d\nTitle: %s\nDate: %s\nParticipants: %s\nSummary: %s\n\n",
			meeting.ID, meeting.Title, meeting.Date.Format("2006-01-02 15:04"), meeting.Participants, meeting.Summary)
	}

	if response == "" {
		response = "Нет запланированных встреч."
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
