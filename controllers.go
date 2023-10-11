package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {

	byteValue, err := json.Marshal(books)
	HandleError(err, "Error Marshalling books into byteValue")
	if string(byteValue) == "null" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No books found in the database"))
		return
	}

	fmt.Println("byteValue String:::", string(byteValue))
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
	w.Header().Set("content-type", "application/json")
}

func getBookById(w http.ResponseWriter, r *http.Request) {

	value := getValueFromPathParams(r, "id")
	id, err := strconv.Atoi(value)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in converting id:" + value + " to integer"))
		return
	}
	fmt.Println("id integer:::", id)

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + string(rune(id)) + " not found"))

	} else {
		jsonBook, _ := json.Marshal(books[index])
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBook)
	}

}

func addBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	HandleError(err, "Error Decoding req body")
	// Handle::Dont add book if some attributes are not present

	books = append(books, book)
	fmt.Println("book:::", book)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Added Successfully"))
}

func updateBookById(w http.ResponseWriter, r *http.Request) {
	value := getValueFromPathParams(r, "id")
	id, err := strconv.Atoi(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in converting id:" + value + " to integer"))
		return
	}

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + value + " not found, so not updating"))

	} else {
		var updatedBook Book
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err == io.EOF {
			// handle if body not found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Body not found in request body"))

		}
		HandleError(err, "Error Decoding req body")

		// handle if body structure not in the form of Book's struct

		books[index] = updatedBook
		w.WriteHeader(http.StatusOK)
		byteUpdatedBook, err := json.Marshal(updatedBook)
		HandleError(err, "Error in Marshalling updatedBook to []byte")

		w.Write(byteUpdatedBook)
	}
}

func deleteBookById(w http.ResponseWriter, r *http.Request) {
	value := getValueFromPathParams(r, "id")
	id, err := strconv.Atoi(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in converting id:" + value + " to integer"))
		return
	}

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + value + " not found, so not deleting"))

	} else {
		deletedBook := books[index]
		books = append(books[:index], books[index+1:]...)
		w.WriteHeader(http.StatusOK)
		byteUpdatedBook, err := json.Marshal(deletedBook)
		HandleError(err, "Error in Marshalling updatedBook to []byte")

		w.Write(byteUpdatedBook)
	}
}
