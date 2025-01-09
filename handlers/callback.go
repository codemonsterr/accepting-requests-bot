package handlers

import (
	"TG_Bot2/utils"
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update, config *utils.Config) {
	if update.CallbackQuery == nil {
		log.Println("CallbackQuery is nil")
		return
	}

	switch update.CallbackQuery.Data {
	case "check_subscription":
		userID := update.CallbackQuery.From.ID
		messageText := handleCheckSubscription(ctx, b, config, userID)

		// Respond to the callback query
		err := utils.AnswerCallbackQuery(ctx, b, update.CallbackQuery.ID)
		if err != nil {
			utils.SendErrorMessage(ctx, b, userID, "Произошла ошибка при подтверждении callback-запроса.")
			return
		}

		// Send the message
		err = utils.SendMessage(ctx, b, userID, messageText)
		if err != nil {
			utils.SendErrorMessage(ctx, b, userID, "Произошла ошибка при отправке сообщения.")
			return
		}
	}
}

// Стоит вынести approveJoinRequest, answerCallbackQuery, sendMessage, sendErrorMessage и вынести в отдельный файл эти функции, чтобы уменьшить код и улучшить читаемость
