package handlers

import (
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
	"log"
)

func handleCheckSubscription(ctx context.Context, b *bot.Bot, config *utils.Config, userID int64) string {
	isSubscribed, err := utils.CheckSubscription(ctx, b, config.Telegram.Channels.TargetChannelID, userID)
	if err != nil {
		log.Printf("Ошибка проверки подписки: %v", err)
		return "Произошла ошибка при проверке подписки. Попробуйте позже."
	}

	if isSubscribed {
		_, approveErr := b.ApproveChatJoinRequest(ctx, &bot.ApproveChatJoinRequestParams{
			ChatID: config.Telegram.Channels.JoinRequestChatID,
			UserID: userID,
		})
		if approveErr != nil {
			log.Printf("Ошибка одобрения заявки пользователя %d: %v", userID, approveErr)

			return "Вы подписаны на канал, но произошла ошибка при одобрении заявки. Попробуйте позже."
		}
		log.Printf("Заявка пользователя %d успешно одобрена", userID)
		return "Вы подписаны на канал! 🎉 Ваша заявка будет одобрена."
	}
	return "Вы не подписаны на канал. Пожалуйста, подпишитесь и попробуйте снова."
}
