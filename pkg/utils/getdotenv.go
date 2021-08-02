package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GoDotEnvVariable(key string) string {
	projectFilePath := os.Getenv("ENV_FILE_PATH")
	// load .env file

	if len(projectFilePath) == 0 {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	} else {
		err := godotenv.Load(projectFilePath +"/.env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	return os.Getenv(key)
}
