package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file:", err)
	}
	fmt.Println(os.Getenv("MYSQL_DSN"))
}
