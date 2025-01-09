package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
)

func ApproveJoinRequest(ctx context.Context, b *bot.Bot, chatID int64, userID int64) {
	_, err := b.ApproveChatJoinRequest(ctx, &bot.ApproveChatJoinRequestParams{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		log.Printf("Ошибка при одобрении заявки пользователя %d: %v", userID, err)
	} else {
		log.Printf("Заявка успешно одобрена для пользователя %d", userID)
	}
}