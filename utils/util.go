package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvOrDefault(envVariable, defaultValue string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	value := os.Getenv(envVariable)
	if value == "" {
		return defaultValue
	}

	return value
}
