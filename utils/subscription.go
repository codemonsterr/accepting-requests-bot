package utils

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
)

func CheckSubscription(ctx context.Context, b *bot.Bot, targetChannelID string, userID int64) (bool, error) {
	member, err := b.GetChatMember(ctx, &bot.GetChatMemberParams{
		ChatID: targetChannelID,
		UserID: userID,
	})
	if err != nil {
		return false, fmt.Errorf("ошибка получения информации о пользователе: %v", err)
	}
	switch member.Type {
	case "member", "administrator", "creator":
		return true, nil
	default:
		return false, nil
	}
}
