package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot struct {
		Token string `yaml:"token"`
	} `yaml:"bot"`
	Channels struct {
		TargetChannelID   string `yaml:"target_channel_id"`
		JoinRequestChatID string `yaml:"join_request_chat_id"`
	} `yaml:"channels"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
