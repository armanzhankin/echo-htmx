package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Host      string
	Port      string
	DB        string
	SecretKey string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}
	return Config{}, nil
}
