package handlers

import (
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

/*
Интерфейсы стоит вынести вот так и передавать их в тайпе
type JoinRequestHandlerFunc func(ctx context.Context, b *bot.Bot, update *models.ChatJoinRequest, config *utils.Config)
*/

// Также стоит потом сделать структуру бот свою у которой будет методы
type UpdateHandler struct {
	JoinRequestHandler   func(ctx context.Context, b *bot.Bot, update *models.ChatJoinRequest, config *utils.Config)
	CallbackQueryHandler func(ctx context.Context, b *bot.Bot, update *models.Update) // Тут тоже нужно передавать конфиг
	MessageHandler       func(ctx context.Context, b *bot.Bot, update *models.Update)
}

func NewUpdateHandler(config *utils.Config) *UpdateHandler { // аргумент конфиг не используется
	return &UpdateHandler{
		JoinRequestHandler:   JoinRequestHandler,
		CallbackQueryHandler: CallbackHandler,
		MessageHandler: func(ctx context.Context, b *bot.Bot, update *models.Update) { // тоже стоит вынести в функцию лучше, или вообще убрать
			if update.Message != nil {
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
	if update.Message != nil { // Возможно стоит поменять на switch
		h.MessageHandler(ctx, b, update)
	}
	if update.ChatJoinRequest != nil {
		h.JoinRequestHandler(ctx, b, update.ChatJoinRequest, config)
	}
	if update.CallbackQuery != nil {
		h.CallbackQueryHandler(ctx, b, update)
	}
}
