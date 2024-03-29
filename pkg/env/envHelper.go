package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(key string, relativePath string) string {
	err := godotenv.Load(relativePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if os.Getenv(key) == "" {
		log.Fatalf("Env key is empty")
	}
	return os.Getenv(key)
}
