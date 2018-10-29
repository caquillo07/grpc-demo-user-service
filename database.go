package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func CreateConnection() (*gorm.DB, error) {
	// Get DB details from the env variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, DBName, password)
	return gorm.Open("postgres", uri)
}
