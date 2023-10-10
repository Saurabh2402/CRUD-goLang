package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var books []Book

func main() {

	// books = append(books, sampleBooks...)
	// for _, book := range books {
	// 	fmt.Printf("book:: %+v\n", book)
	// }

	r := mux.NewRouter()

	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBookById).Methods("GET")
	r.HandleFunc("/book", addBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBookById).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBookById).Methods("DELETE")
	http.Handle("/", r)

	godotenv.Load()
	port := GetEnvVar("PORT")

	fmt.Printf("Server starting at:  http://localhost:%s\n\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server Failed to start, error:", err)
	}
	// Code below this point will continue executing
	fmt.Println("Server started on", port)
}
