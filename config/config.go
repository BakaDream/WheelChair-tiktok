package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file:", err)
	}
}
