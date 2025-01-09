package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
)

func sendErrorMessage(ctx context.Context, b *bot.Bot, userID int64, message string) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   message,
	})
	if err != nil {
		log.Printf("Ошибка отправки сообщения пользователю %d: %v", userID, err)
	}
}
