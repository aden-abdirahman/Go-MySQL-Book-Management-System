package models

import (
	"github.com/aden-abdirahman/Go-MySQL-Book-Management-System/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// creating our book struct
type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init func that opens our mysql connection and gets our database
func init() {
	config.Connect()
	db = config.GetDB()
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. AutoMigrate creates database foreign key constraints automatically
	db.AutoMigrate(&Book{})
}

// create book method that checks to see if primary key is blank then inserts it into the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// getting all the books in the db
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// func that returns a specific book by filtering the book records by id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book

	db := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

// func that deletes books by id
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
