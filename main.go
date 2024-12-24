package main

import (
	"TG_Bot2/handlers"
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Загрузка конфигурации
	config, err := utils.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Создание экземпляра обработчика обновлений
	updateHandler := handlers.NewUpdateHandler(config)

	// Создаём контекст с поддержкой завершения через сигнал
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Создание бота
	opts := []bot.Option{
		bot.WithDefaultHandler(func(ctx context.Context, b *bot.Bot, update *models.Update) {
			updateHandler.HandleUpdate(ctx, b, update, config)
		}),
	}

	// Создаём экземпляр бота
	b, err := bot.New(config.Bot.Token, opts...)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	// Запуск бота
	b.Start(ctx)
}
