package config

import (
	"log"
	"os"
)

type Config struct {
	JWT_SECRET  string
	APP_NAME    string
	APP_NODE    string
	APP_PORT    string
	APP_ORIGINS string
	DB_URL      string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		APP_NAME:    os.Getenv("APP_NAME"),
		APP_NODE:    os.Getenv("APP_NODE"),
		APP_PORT:    os.Getenv("APP_PORT"),
		APP_ORIGINS: os.Getenv("APP_ORIGINS"),
		DB_URL:      os.Getenv("DATABASE_URL"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}

	if AppConfig.APP_NAME == "" {
		AppConfig.APP_NAME = "GOLANG API"
	}

	if AppConfig.APP_NODE == "" {
		AppConfig.APP_NODE = "development"
	}

	if AppConfig.APP_PORT == "" {
		AppConfig.APP_PORT = "3000"
	}

	if AppConfig.APP_ORIGINS == "" {
		AppConfig.APP_ORIGINS = "*"
	}

	if AppConfig.JWT_SECRET == "" {
		AppConfig.JWT_SECRET = "default_secret_key"
	}

	if AppConfig.DB_URL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}
}
