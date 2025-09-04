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

	// Set trusted IPs
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	// Registers routes
	router.GET("/user/:id", getUserProfile)
	router.POST("/user", registerUserProfile)
	router.PATCH("/user/:id", updateUserProfile)
	router.DELETE("/user/:id", deleteUserProfile)

	router.Run("localhost:" + port)
	log.Println("Server is listening on localhost:", port)
}
