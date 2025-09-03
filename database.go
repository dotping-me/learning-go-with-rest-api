package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global pointer to database instance

func initDB(dbName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database!")
	}

	err = DB.AutoMigrate(&UserProfile{}) // Migrates data schemas to match models
	if err != nil {
		log.Fatal("Failed to mitigate Database!")
	}
}
