package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func HandleError(err error, errorMessage string) {
	if err != nil {
		fmt.Printf("Error:: %s, %v", errorMessage, err)
	}
}
func GetEnvVar(key string) string {
	err := godotenv.Load(".env")
	HandleError(err, "Error loading .env file")

	return os.Getenv(key)
}
