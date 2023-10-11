package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

func findBookIndexById(id int) int {

	for ind, book := range books {
		if id == book.Id {
			fmt.Println("Book Found!")
			return ind
		}
	}
	fmt.Println("Book Not Found, returning -1")

	return -1
}

func getValueFromPathParams(r *http.Request, key string) string {
	vars := mux.Vars(r)
	value := vars[key]
	fmt.Println("key:", key)
	fmt.Println("value:", value)
	return value
}
