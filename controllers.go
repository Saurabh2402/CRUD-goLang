package main

import (
	"fmt"
	"net/http"
)

func getAllBooks(http.ResponseWriter, *http.Request) {
	fmt.Println("getAllBooks()")
}

func getBookById(http.ResponseWriter, *http.Request) {
	fmt.Println("getBookById()")

}

func addBook(http.ResponseWriter, *http.Request) {
	fmt.Println("addBook()")

}

func updateBookById(http.ResponseWriter, *http.Request) {
	fmt.Println("updateBookById()")

}

func deleteBookById(http.ResponseWriter, *http.Request) {
	fmt.Println("deleteBookById()")

}
