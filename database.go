package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	// Get DB details from the env variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	uri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, DBName)
	return gorm.Open("postgres", uri)
}
