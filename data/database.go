/*

Database Initialisation
- Connection is established globally (i.e. No return values)

*/

package data

import (
	"log"

	"github.com/dotping-me/learning-go-with-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnetToDB(dsn string) {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database: ", err)
	}

	// Migrates data schemas to match models
	err = DB.AutoMigrate(&models.UserProfile{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to mitigate Database: ", err)
	}
}
