package utils

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v9"
	"gopkg.in/yaml.v3"
)

type TelegramConfig struct {
	Bot struct {
		Token string `yaml:"token" env:"TELEGRAM_BOT_TOKEN"`
	} `yaml:"bot"`
	Channels struct {
		TargetChannelID   string `yaml:"target_channel_id" env:"TARGET_CHANNEL_ID"`
		JoinRequestChatID string `yaml:"join_request_chat_id" env:"JOIN_REQUEST_CHAT_ID"`
	} `yaml:"channels"`
}

type Config struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

func LoadConfig(path string) (*Config, error) {
	// Загружаем YAML-файл
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла конфигурации: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML: %w", err)
	}

	// Переопределяем значениями из переменных окружения
	if err := env.Parse(&config.Telegram); err != nil {
		return nil, fmt.Errorf("ошибка парсинга переменных окружения: %w", err)
	}

	return &config, nil
}
