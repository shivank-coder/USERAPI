package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// for loading the env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")

	}
}

// fatching the keys from .env file by using function os.Getenv file
func GetEnv(key string) string {
	return os.Getenv(key)
}
