package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL          string
	JWTSecret      string
	Port           string
	AppLogFile     string
	ServiceLogFile string
	HandlerLogFile string
}

func LoadConfig() *Config {
	// Godotenv.Load ищет .env в текущей директории или выше.
	// Если main.go находится в cmd/web, а .env в корне, то путь будет "../../.env"
	// Или используйте os.Getenv без godotenv.Load, если переменные окружения уже заданы в системе.
	err := godotenv.Load("../../.env") // Путь относительно cmd/web/main.go
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBURL:          os.Getenv("DB_URL"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		Port:           os.Getenv("PORT"),
		AppLogFile:     os.Getenv("APP_LOG_FILE"),
		ServiceLogFile: os.Getenv("SERVICE_LOG_FILE"),
		HandlerLogFile: os.Getenv("HANDLER_LOG_FILE"),
	}
}
