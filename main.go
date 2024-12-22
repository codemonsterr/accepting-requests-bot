package main

import (
	"TG_Bot2/handlers"
	"TG_Bot2/utils"
	"context"
	"github.com/go-telegram/bot"
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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.MessageHandler),
	}

	b, err := bot.New(config.Bot.Token, opts...)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	b.Start(ctx)
}
