package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global pointer to database instance

func initDB(dbURI string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database!")
	}

	// Migrates data schemas to match models
	err = DB.AutoMigrate(&UserProfile{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("Failed to mitigate Database!")
	}
}
