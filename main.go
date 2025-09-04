package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Loads .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialises Database
	initDB(os.Getenv("DB_NAME"))

	// Starts Gin router
	port := os.Getenv("PORT")
	router := gin.Default()

	// Registers routes
	router.GET("/user/profile", getUserProfile)
	router.POST("/user/profile", registerUserProfile)
	router.PATCH("/user/profile", updateUserProfile)

	router.Run("localhost:" + port)
	log.Println("Server is listening on localhost:", port)
}
