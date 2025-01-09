package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-telegram/bot"
)

const ( // Отлично
	Member        = "member"
	Administrator = "administrator"
	Creator       = "creator"
)

func CheckSubscription(ctx context.Context, b *bot.Bot, targetChannelID string, userID int64) (bool, error) {
	if targetChannelID == "" {
		return false, errors.New("targetChannelID не должен быть пустым")
	}

	member, err := b.GetChatMember(ctx, &bot.GetChatMemberParams{
		ChatID: targetChannelID,
		UserID: userID,
	})
	if err != nil {
		return false, fmt.Errorf("ошибка получения информации о пользователе (userID: %d, channelID: %s): %v", userID, targetChannelID, err)
	}
	switch member.Type {
	case Member, Administrator, Creator:
		return true, nil
	default:
		return false, nil
	}
}


