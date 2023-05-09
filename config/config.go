package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	SmtpUsername string
	SmtpPassword string

	UserEmail string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &AppConfig{
		SmtpUsername: os.Getenv("SMTP_USERNAME"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
		UserEmail:    os.Getenv("USER_EMAIL"),
	}
}
