package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// Стоит добавить проверку update.CallbackQuery на nil
	// Конфиг нужно инициализироваться один раз и передаваться в аргументах функции, чтобы избежать дублирование кода, и лучше переписать конфиг с этой библиотекой https://github.com/caarlos0/env
	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		log.Printf("Ошибка загрузки конфигурации: %v", err)
		return
	}

	callbackQueryID := update.CallbackQuery.ID
	// Возможно стоит поменять на switch конструкцию
	if update.CallbackQuery.Data == "check_subscription" {
		// Вынести тело условия в отдельную функцию для более удобной читаемости
		userID := update.CallbackQuery.From.ID

		isSubscribed, err := utils.CheckSubscription(ctx, b, config.Channels.TargetChannelID, userID)
		if err != nil {
			log.Printf("Ошибка проверки подписки: %v", err)
			return
		}

		messageText := "Вы не подписаны на канал. Пожалуйста, подпишитесь и попробуйте снова."
		if isSubscribed {
			messageText = "Вы подписаны на канал! 🎉 Ваша заявка будет одобрена."

			// Одобрение заявки автоматически
			/*Стоит вынести approveJoinRequest, answerCallbackQuery, sendMessage, sendErrorMessage и вынести в отдельный файл эти функции, чтобы уменьшить код и улучшить читаемость
			Пример функции

			func sendErrorMessage(ctx context.Context, b *bot.Bot, chatID int64, text string) {
				err := sendMessage(ctx, b, chatID, text)
				if err != nil {
					log.Printf("Ошибка отправки сообщения об ошибке пользователю %d: %v", chatID, err)
				}
			}*/
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

		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text:   messageText,
		})
		if err != nil {
			log.Printf("Ошибка отправки сообщения: %v", err)
		}
	}
}
