package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func JoinRequestHandler(ctx context.Context, b *bot.Bot, request *models.ChatJoinRequest, config *utils.Config /*Молодец, также в других надо делать*/) {
	// Стоит добавить проверку request и также request.Chat и request.From
	chatID := request.Chat.ID
	userID := request.From.ID

	isSubscribed, err := utils.CheckSubscription(ctx, b, config.Channels.TargetChannelID, userID)
	if err != nil {
		log.Printf("Ошибка проверки подписки: %v", err)
		// Стоит использовать sendErrorMessage описанную в другом файле
		return
	}

	if isSubscribed {
		// Также вынести функции в отдельный файл
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
		sendCheckButton(ctx, b, *config, userID)
	}
}

func sendCheckButton(ctx context.Context, b *bot.Bot, config utils.Config, userID int64) {
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
