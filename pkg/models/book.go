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
