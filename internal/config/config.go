package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BotToken string `yaml:"bot_token"`
	DbUrl    string `yaml:"db_url"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	yamlFile, err := os.ReadFile("./conf.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
