package main

import (
	"fmt"
	"net/http"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllBooks()")
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getBookById()")

}

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addBook()")

}

func updateBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateBookById()")

}

func deleteBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteBookById()")

}
