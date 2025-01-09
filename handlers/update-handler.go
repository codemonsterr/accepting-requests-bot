package handlers

import (
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

type JoinRequestHandlerFunc func(ctx context.Context, b *bot.Bot, update *models.ChatJoinRequest, config *utils.Config)

type CallbackQueryHandlerFunc func(ctx context.Context, b *bot.Bot, update *models.Update, config *utils.Config)

type UpdateHandler struct {
	JoinRequestHandler   JoinRequestHandlerFunc
	CallbackQueryHandler CallbackQueryHandlerFunc
}

func NewUpdateHandler() *UpdateHandler {
	return &UpdateHandler{
		JoinRequestHandler:   JoinRequestHandler,
		CallbackQueryHandler: CallbackHandler,
	}
}

func (h *UpdateHandler) HandleUpdate(ctx context.Context, b *bot.Bot, update *models.Update, config *utils.Config) {
	switch {
	case update.ChatJoinRequest != nil:
		h.JoinRequestHandler(ctx, b, update.ChatJoinRequest, config)
	case update.CallbackQuery != nil:
		h.CallbackQueryHandler(ctx, b, update, config)
	default:
		log.Println("Ошибка выполнения проверки: неизвестный тип обновления")
	}
}
