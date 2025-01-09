package utils

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
)

const (
	tgLink = "https://t.me/" // хуй его знает как можно было еще назвать это
)

func GetChannelLink(ctx context.Context, b *bot.Bot, targetChannelID string) (string, error) {
	chat, err := b.GetChat(ctx, &bot.GetChatParams{
		ChatID: targetChannelID,
	})
	if err != nil {
		return "", fmt.Errorf("не удалось получить данные о канале: %w", err)
	}

	if chat.Username != "" {
		return fmt.Sprintf("%s%s", tgLink, chat.Username), nil
	}
	return "", fmt.Errorf("у канала нет публичного имени (это приватный канал)") // стоит немного отформатировать строки ошибок (необязательно)
}
