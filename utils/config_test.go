package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name          string
		setupFile     func() string
		setupEnv      func()
		expectedError string
	}{
		{
			name: "successful config load",
			setupFile: func() string {
				// Создаем временный файл с YAML конфигурацией
				fileName := "test_config.yaml"
				yamlContent := `
telegram:
  bot:
    token: "dummy_token"
  channels:
    target_channel_id: "target_channel_id"
    join_request_chat_id: "join_request_chat_id"`
				err := os.WriteFile(fileName, []byte(yamlContent), 0644)
				require.NoError(t, err)
				return fileName
			},
			setupEnv: func() {
				// Не задаем переменные окружения для теста (используется файл)
			},
			expectedError: "",
		},
		{
			name: "file open error",
			setupFile: func() string {
				// Возвращаем несуществующий файл
				return "non_existent_file.yaml"
			},
			setupEnv:      func() {},
			expectedError: "ошибка открытия файла конфигурации",
		},
		{
			name: "yaml parse error",
			setupFile: func() string {
				// Создаем файл с ошибочным YAML
				fileName := "invalid_config.yaml"
				yamlContent := `
telegram:
  bot:
    token: dummy_token
  channels:
    target_channel_id: target_channel_id
    join_request_chat_id: join_request_chat_id`
				err := os.WriteFile(fileName, []byte(yamlContent), 0644)
				require.NoError(t, err)
				return fileName
			},
			setupEnv:      func() {},
			expectedError: "ошибка парсинга YAML",
		},
		{
			name: "env parse error",
			setupFile: func() string {
				// Создаем корректный файл, но ошибочные переменные окружения
				fileName := "valid_config.yaml"
				yamlContent := `
telegram:
  bot:
    token: "dummy_token"
  channels:
    target_channel_id: "target_channel_id"
    join_request_chat_id: "join_request_chat_id"`
				err := os.WriteFile(fileName, []byte(yamlContent), 0644)
				require.NoError(t, err)
				return fileName
			},
			setupEnv: func() {
				// Устанавливаем неверные переменные окружения для теста
				_ = os.Setenv("TELEGRAM_BOT_TOKEN", "")
			},
			expectedError: "ошибка парсинга переменных окружения",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Подготавливаем файл и переменные окружения
			fileName := tt.setupFile()
			tt.setupEnv()

			// Пытаемся загрузить конфигурацию
			config, err := LoadConfig(fileName)
			if tt.expectedError == "" {
				// Если ошибка не ожидалась, проверяем успешную загрузку
				require.NoError(t, err)
				assert.Equal(t, "dummy_token", config.Telegram.Bot.Token)
				assert.Equal(t, "target_channel_id", config.Telegram.Channels.TargetChannelID)
				assert.Equal(t, "join_request_chat_id", config.Telegram.Channels.JoinRequestChatID)
			} else {
				// Проверяем, что ошибка соответствует ожидаемой
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}

			// Удаляем файл после теста
			_ = os.Remove(fileName)

			// Чистим переменные окружения после теста
			_ = os.Unsetenv("TELEGRAM_BOT_TOKEN")
			_ = os.Unsetenv("TARGET_CHANNEL_ID")
			_ = os.Unsetenv("JOIN_REQUEST_CHAT_ID")
		})
	}
}
