package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// ConnectDB establishes a connection to the database using the provided connection string
func ConnectDB() {
	d, err := gorm.Open("mysql", "root:Password@tcp(localhost:3306)/mydatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Error", err)
		panic("failed to connect database")

	}
	fmt.Println("Connected to database", d)
	db = d
}

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection
