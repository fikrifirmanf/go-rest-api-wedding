package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	env := make(chan string, 1)
	// load .env file
	envFile := godotenv.Load(".env")
	if envFile != nil {
		log.Fatalf("Error loading .env file")
	} else {
		if os.Getenv("ENVIRONMENT") != "production" {
			env <- os.Getenv(key)
		} else {
			env <- os.Getenv(key)
		}
	}

	return <-env
}
