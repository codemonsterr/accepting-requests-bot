package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func MessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
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
