package handlers

import (
	"log"
	"schedulertgbot/db"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAddSummary(bot *tgbotapi.BotAPI, message *tgbotapi.Message, repo *db.MeetingRepository) {
	args := message.CommandArguments()
	parts := strings.SplitN(args, ";", 2)
	if len(parts) != 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный формат. Используйте: /add_summary meeting_id; summary")
		bot.Send(msg)
		return
	}

	meetingID, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный формат ID встречи.")
		bot.Send(msg)
		return
	}

	summary := strings.TrimSpace(parts[1])
	err = repo.AddSummary(meetingID, summary)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Итоги встречи успешно добавлены!")
	bot.Send(msg)
}
