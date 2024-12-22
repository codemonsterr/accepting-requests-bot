package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		log.Printf("Ошибка загрузки конфигурации: %v", err)
		return
	}

	callbackQueryID := update.CallbackQuery.ID
	if update.CallbackQuery.Data == "check_subscription" {
		userID := update.CallbackQuery.From.ID

		isSubscribed, err := utils.CheckSubscription(ctx, b, config.Channels.TargetChannelID, userID)
		if err != nil {
			log.Printf("Ошибка проверки подписки: %v", err)
			b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: callbackQueryID,
				Text:            "Произошла ошибка при проверке подписки. Попробуйте позже.",
				ShowAlert:       true,
			})
			return
		}

		messageText := "Вы не подписаны на канал. Пожалуйста, подпишитесь и попробуйте снова."
		if isSubscribed {
			messageText = "Вы подписаны на канал! 🎉 Ваша заявка будет одобрена."

			// Одобрение заявки автоматически
			_, approveErr := b.ApproveChatJoinRequest(ctx, &bot.ApproveChatJoinRequestParams{
				ChatID: config.Channels.JoinRequestChatID,
				UserID: userID,
			})
			if approveErr != nil {
				log.Printf("Ошибка одобрения заявки пользователя %d: %v", userID, approveErr)
				messageText = "Вы подписаны на канал, но произошла ошибка при одобрении заявки. Попробуйте позже."
			} else {
				log.Printf("Заявка пользователя %d успешно одобрена", userID)
			}
		}

		_, err = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: callbackQueryID,
		})
		if err != nil {
			log.Printf("Ошибка подтверждения callback-запроса: %v", err)
		}

		// Уведомляем пользователя
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   messageText,
		})
		if err != nil {
			log.Printf("Ошибка отправки сообщения: %v", err)
		}

	}

}
