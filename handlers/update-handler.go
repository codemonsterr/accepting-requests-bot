package handlers

import (
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

type UpdateHandler struct {
	JoinRequestHandler   func(ctx context.Context, b *bot.Bot, update *models.ChatJoinRequest, config *utils.Config)
	CallbackQueryHandler func(ctx context.Context, b *bot.Bot, update *models.Update)
	MessageHandler       func(ctx context.Context, b *bot.Bot, update *models.Update)
}

func NewUpdateHandler(config *utils.Config) *UpdateHandler {
	return &UpdateHandler{
		JoinRequestHandler:   JoinRequestHandler, // передаем конфиг в обработчик
		CallbackQueryHandler: CallbackHandler,
		MessageHandler: func(ctx context.Context, b *bot.Bot, update *models.Update) {
			// Обрабатываем сообщения
			if update.Message != nil {
				// следует обрабатывать ошибки
				_, err := b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   update.Message.Text,
				})
				if err != nil {
					// Логируем ошибку отправки сообщения
					log.Printf("Ошибка отправки сообщения: %v", err)
				}
			}
		},
	}
}

func (h *UpdateHandler) HandleUpdate(ctx context.Context, b *bot.Bot, update *models.Update, config *utils.Config) {
	if update.Message != nil {
		h.MessageHandler(ctx, b, update)
	}
	if update.ChatJoinRequest != nil {
		h.JoinRequestHandler(ctx, b, update.ChatJoinRequest, config)
	}
	if update.CallbackQuery != nil {
		h.CallbackQueryHandler(ctx, b, update)
	}
}
