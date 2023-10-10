package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {

	byteValue, err := json.Marshal(books)
	HandleError(err, "Error Marshalling books into byteValue")

	fmt.Println("byteValue String:::", string(byteValue))
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
	w.Header().Set("content-type", "application/json")
}

func getBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	fmt.Println("idStr:::", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in converting id :" + idStr + " to integer"))
		return
	}

	// HandleError(err, "Error in converting id :"+idStr+" to integer")
	fmt.Println("id integer:::", id)

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + idStr + " not found"))

	} else {
		jsonBook, _ := json.Marshal(books[index])
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBook)
	}

}

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addBook()")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	HandleError(err, "Error Decoding req body")

	books = append(books, book)
	fmt.Println("book:::", book)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Added Successfully"))
}

func updateBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updateBookById()")
	vars := mux.Vars(r)
	idStr := vars["id"]
	// fmt.Println("idStr:::", idStr)
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:::", id)

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + idStr + " not found, so not updating"))

	} else {
		var updatedBook Book
		err := json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			// handle if body not found
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
	fmt.Println("deleteBookById()")
	vars := mux.Vars(r)
	idStr := vars["id"]
	// fmt.Println("idStr:::", idStr)
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:::", id)

	index := findBookIndexById(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book with id::" + idStr + " not found, so not deleting"))

	} else {
		deletedBook := books[index]
		books = append(books[:index], books[index+1:]...)
		w.WriteHeader(http.StatusOK)
		byteUpdatedBook, err := json.Marshal(deletedBook)
		HandleError(err, "Error in Marshalling updatedBook to []byte")

		w.Write(byteUpdatedBook)
	}
}
