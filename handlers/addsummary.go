package handlers

import (
	"database/sql"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAddSummary(bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *sql.DB) {
	args := message.CommandArguments()
	parts := strings.SplitN(args, ";", 2)
	if len(parts) != 2 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неверный формат. Используйте: /add_summary meeting_id; summary")
		bot.Send(msg)
		return
	}

	meetingID := strings.TrimSpace(parts[0])
	summary := strings.TrimSpace(parts[1])

	_, err := db.Exec("UPDATE meetings SET summary = $1 WHERE id = $2", summary, meetingID)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Итоги встречи успешно добавлены!")
	bot.Send(msg)
}
