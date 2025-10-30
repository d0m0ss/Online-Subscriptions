package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func (ac *AppConfig) GetMainSettings() map[string]string {
	return map[string]string{
		"SERVER_HOST":  os.Getenv("SERVER_HOST"),
		"SERVER_PORT":  os.Getenv("SERVER_PORT"),
		"CORS_ORIGINS": os.Getenv("CORS_ORIGINS"),
		"DEBUG":        os.Getenv("DEBUG"),
	}
}

func (ac *AppConfig) GetDB() string {
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	prefix := "postgresql"

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", prefix, dbUser, dbPassword, dbHost, dbPort, dbName)
}
