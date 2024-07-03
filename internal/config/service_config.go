package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type ServiceConfig struct {
	CheckerBotToken string   `yaml:"checker_bot_token" env-required:"true"`
	ChatId          int64    `yaml:"chat_td" env-required:"true"`
	Sites           []string `yaml:"sites" env-required:"true"`
}

func MustLoad() *ServiceConfig {
	configPath := os.Getenv("SERVICE_CONFIG_PATH")
	if len(configPath) == 0 {
		log.Fatal("SERVICE_CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file '%s' does not exist", configPath)
	}

	var cfg ServiceConfig
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can't read config: %s", err)
	}

	return &cfg
}
