package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aden-abdirahman/Go-MySQL-Book-Management-System/pkg/models"
	"github.com/aden-abdirahman/Go-MySQL-Book-Management-System/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// each function has a corrosponding function in our models folder
func GetBook(w http.ResponseWriter, r *http.Request) {
	// storing our list of books inside our newbooks var
	newBooks := models.GetAllBooks()
	// setting res to our newbooks list, json encoded, and writing it to our response
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// storing our req variables in vars
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// parsing our book id to int64
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// running our getbookbyid function from our models folder and storing the book in bookdetails
	bookDetails, _ := models.GetBookById(ID)

	// sending our json encoded book to user
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// setting createBook to a models.Book struct
	CreateBook := &models.Book{}
	// parsing the response from json so the database can understand
	utils.ParseBody(r, CreateBook)
	// running the create book function on our create book var
	b := CreateBook.CreateBook()
	// sending our marshaled book to the user
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(ID)
	// sending the deleted book to user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// setting updatebook to models.book
	var updateBook = &models.Book{}
	// parsing req
	utils.ParseBody(r, updateBook)
	// setting our variables
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	// converting book id to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// searching for book in our db
	bookDetails, db := models.GetBookById(ID)
	// replacing the info of the book in the database with the new info from the req
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
