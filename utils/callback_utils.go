package utils

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
)

func AnswerCallbackQuery(ctx context.Context, b *bot.Bot, callbackQueryID string) error {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: callbackQueryID,
	})
	if err != nil {
		log.Printf("Ошибка подтверждения callback-запроса: %v", err)
	}
	return err
}

// Sends a text message
func SendMessage(ctx context.Context, b *bot.Bot, chatID int64, text string) error {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
	return err
}

// Sends an error message to the user
func SendErrorMessage(ctx context.Context, b *bot.Bot, chatID int64, text string) {
	err := SendMessage(ctx, b, chatID, text)
	if err != nil {
		log.Printf("Ошибка отправки сообщения об ошибке: %v", err)
	}
}
