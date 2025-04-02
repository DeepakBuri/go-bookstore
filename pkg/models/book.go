package models

import (
	"github.com/deepakburi/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Book represents the structure of a book in the bookstore

func (book *Book) CreateBook() *Book {
	db.NewRecord(book) // set to false
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(bookID int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", bookID).Find(&book)
	return &book, db
}

func DeleteBook(bookID int64) Book {
	var book Book
	db.Where("ID=?", bookID).Delete(book)
	return book
}
