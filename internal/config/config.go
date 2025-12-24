package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DbHost     string
	DbName     string
	DbPort     string
	DbUser     string
	DbPassword string

	JwtSecret string

	WebhookURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port: os.Getenv("PORT"),

		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),

		JwtSecret:  os.Getenv("JWT_SECRET"),
		WebhookURL: os.Getenv("WEBHOOK_URL"),
	}
}
