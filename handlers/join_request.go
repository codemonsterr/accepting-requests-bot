package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Обработчик заявки на вступление в чат
func JoinRequestHandler(ctx context.Context, b *bot.Bot, request *models.ChatJoinRequest, config *utils.Config) {

	if request.Chat.ID == 0 || request.From.ID == 0 {
		// Обработка ошибки, если ID равен 0
		sendErrorMessage(ctx, b, request.From.ID, "Ошибка: request.Chat.ID или request.From.ID равны 0.")
		return
	}

	chatID := request.Chat.ID
	userID := request.From.ID

	// Проверка подписки
	isSubscribed, err := utils.CheckSubscription(ctx, b, config.Telegram.Channels.TargetChannelID, userID)
	if err != nil {
		log.Printf("Ошибка проверки подписки: %v", err)
		sendErrorMessage(ctx, b, userID, "Произошла ошибка при проверке подписки.")
		return
	}

	// Обработка заявки в зависимости от подписки
	if isSubscribed {
		ApproveJoinRequest(ctx, b, chatID, userID)
	} else {
		sendCheckButton(ctx, b, *config, userID)
	}
}

// Функция для отправки кнопки с проверкой подписки
func sendCheckButton(ctx context.Context, b *bot.Bot, config utils.Config, userID int64) {
	// Получаем ссылку на канал
	channelLink, err := utils.GetChannelLink(ctx, b, config.Telegram.Channels.TargetChannelID)
	if err != nil {
		log.Printf("Ошибка получения ссылки на канал: %v", err)
		sendErrorMessage(ctx, b, userID, "Ошибка получения ссылки на канал.")
		return
	}

	// Отправляем сообщение с кнопками
	_, sendErr := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   "Чтобы ваша заявка была одобрена, подпишитесь на наш канал!",
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: [][]models.InlineKeyboardButton{
				{
					{
						Text: "Перейти на канал",
						URL:  channelLink,
					},
				},
				{
					{
						Text:         "Проверить подписку",
						CallbackData: "check_subscription",
					},
				},
			},
		},
	})
	if sendErr != nil {
		log.Printf("Ошибка отправки сообщения: %v", sendErr)
	}
}
