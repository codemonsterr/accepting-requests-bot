package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Handler должен отвечать только за логику взаимодействия с реквестом
func JoinRequestHandler(ctx context.Context, b *bot.Bot, request *models.ChatJoinRequest) {
	// конфиг нужно передавать в аргументах функции
	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		log.Printf("Ошибка загрузки конфигурации: %v", err)
		return
	}

	chatID := request.Chat.ID
	userID := request.From.ID

	// одинаковые ошибки должны одинаково обрабатываться
	isSubscribed, err := utils.CheckSubscription(ctx, b, config.Channels.TargetChannelID, userID)
	if err != nil {
		log.Printf("Ошибка проверки подписки: %v", err)
		return
	}

	if isSubscribed {
		_, err := b.ApproveChatJoinRequest(ctx, &bot.ApproveChatJoinRequestParams{
			ChatID: chatID,
			UserID: userID,
		})
		if err != nil {
			log.Printf("Ошибка при одобрении заявки пользователя %d: %v", userID, err)
		} else {
			log.Printf("Заявка успешно одобрена %d", userID)
		}
	} else {
		// получение ссылки следует вынести в отдельную функцию
		channelLink, _ := utils.GetChannelLink(ctx, b, config.Channels.TargetChannelID)
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
}
