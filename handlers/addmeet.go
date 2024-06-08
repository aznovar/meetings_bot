package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"schedulertgbot/db"
	"strings"
	"time"
)

func HandleAddMeeting(bot *tgbotapi.BotAPI, message *tgbotapi.Message, repo *db.MeetingRepository) {
	args := message.CommandArguments()
	parts := strings.SplitN(args, ";", 3)
	if len(parts) != 3 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный формат. Используйте: /add_meeting title; date; participants")
		bot.Send(msg)
		return
	}

	title := strings.TrimSpace(parts[0])
	dateStr := strings.TrimSpace(parts[1])
	participants := strings.TrimSpace(parts[2])

	date, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный формат даты. Используйте: YYYY-MM-DD HH:MM")
		bot.Send(msg)
		return
	}

	err = repo.AddMeeting(title, date, participants)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Встреча успешно добавлена!")
	bot.Send(msg)
}
