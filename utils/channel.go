package utils

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
)

func GetChannelLink(ctx context.Context, b *bot.Bot, targetChannelID string) (string, error) {
	chat, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: targetChannelID,
	})
	if err != nil {
		return "", fmt.Errorf("не удалось получить данные о канале: %w", err)
	}

	if chat.Username != "" {
		return "https://t.me/" + chat.Username, nil
	}
	return "", fmt.Errorf("у канала нет публичного имени (это приватный канал)") // стоит немного отформатировать строки ошибок (необязательно)
}
