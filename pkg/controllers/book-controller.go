package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nitin1chandani/Go-proj-book-mgmt/pkg/models"
	"github.com/nitin1chandani/Go-proj-book-mgmt/pkg/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	if err := utils.ParseBody(r, CreateBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	CreatedBook := CreateBook.CreateBook()
	res, err := json.Marshal(CreatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Book Id", http.StatusBadRequest)
		return
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		http.Error(w, "Invalid Body", http.StatusInternalServerError)
		return
	}
	UpdatedBook := book.CreateBook()
	res, err := json.Marshal(UpdatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid Book Id", http.StatusBadRequest)
		return
	}

	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
