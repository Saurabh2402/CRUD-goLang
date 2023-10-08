package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hello World!!!")

	r := mux.NewRouter()

	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBookById).Methods("GET")
	r.HandleFunc("/book", addBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBookById).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBookById).Methods("DELETE")
	http.Handle("/", r)

	godotenv.Load()
	port := GetEnvVar("PORT")
	log.Fatal(http.ListenAndServe(port, nil))

}
