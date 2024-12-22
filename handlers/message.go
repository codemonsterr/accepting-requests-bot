package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// хендлер перехватывает все сообщения, а назван MessageHandler
// Стоит создать структуру UpdateHandler в которой будут MessageHandler JoinRequestHandler CallbackQueryHandler и функцию NewUpdateHandler которая будет возвращать UpdateHandler и принимать структуры описанные ранее
func MessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		// следует обрабатывать ошибки
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
	if update.ChatJoinRequest != nil {
		JoinRequestHandler(ctx, b, update.ChatJoinRequest)
	}
	if update.CallbackQuery != nil {
		CallbackHandler(ctx, b, update)
	}

}
