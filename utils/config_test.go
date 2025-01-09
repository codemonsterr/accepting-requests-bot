package utils

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	configContent := `
telegram:
  bot:
    token: "default_token"
  channels:
    target_channel_id: "default_target_channel_id"
    join_request_chat_id: "default_join_request_chat_id"
`
	configFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(configFile.Name())

	if _, err := configFile.Write([]byte(configContent)); err != nil {
		t.Fatalf("Failed to write to temp config file: %v", err)
	}
	if err := configFile.Close(); err != nil {
		t.Fatalf("Failed to close temp config file: %v", err)
	}

	// Set environment variables to override config values
	os.Setenv("TELEGRAM_BOT_TOKEN", "env_token")
	os.Setenv("TARGET_CHANNEL_ID", "env_target_channel_id")
	os.Setenv("JOIN_REQUEST_CHAT_ID", "env_join_request_chat_id")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")
	defer os.Unsetenv("TARGET_CHANNEL_ID")
	defer os.Unsetenv("JOIN_REQUEST_CHAT_ID")

	// Load the configuration
	config, err := LoadConfig(configFile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify that the configuration values are as expected
	if config.Telegram.Bot.Token != "env_token" {
		t.Errorf("Expected Telegram.Bot.Token to be 'env_token', got '%s'", config.Telegram.Bot.Token)
	}
	if config.Telegram.Channels.TargetChannelID != "env_target_channel_id" {
		t.Errorf("Expected Telegram.Channels.TargetChannelID to be 'env_target_channel_id', got '%s'", config.Telegram.Channels.TargetChannelID)
	}
	if config.Telegram.Channels.JoinRequestChatID != "env_join_request_chat_id" {
		t.Errorf("Expected Telegram.Channels.JoinRequestChatID to be 'env_join_request_chat_id', got '%s'", config.Telegram.Channels.JoinRequestChatID)
	}
}
