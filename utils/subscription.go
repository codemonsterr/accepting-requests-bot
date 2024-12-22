package utils

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
)

// Стоит возвращать более понятные ошибки
func CheckSubscription(ctx context.Context, b *bot.Bot, targetChannelID string, userID int64) (bool, error) {
	member, err := b.GetChatMember(ctx, &bot.GetChatMemberParams{
		ChatID: targetChannelID, // стоит добавить проверку на targetChannelId что он не пустой
		UserID: userID,
	})
	if err != nil {
		return false, fmt.Errorf("ошибка получения информации о пользователе: %v", err)
	}
	switch member.Type {
	// стоит использовать константы в switch
	case "member", "administrator", "creator":
		return true, nil
	default:
		return false, nil
	}
}
