package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariables (key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	result := os.Getenv(key)

	return result
}
