/*
Next Steps:
- Switch to Postgres
- Organise Project Structure
- Implement Users, Posts, Comments and so on
- Dockerise everything!!!

- Look into MVC Structure and Frontend Integration for later
*/

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

	initDB() // Initialises Database

	// Starts Gin router
	port := os.Getenv("PORT")
	router := gin.Default()

	// Set trusted IPs
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	// Using JWT auth
	jwtMiddleware := initJWTMiddleware()
	router.POST("/login", jwtMiddleware.LoginHandler)
	router.POST("/user", registerUserProfile)

	// Groups similar routes together and wraps them with JWT auth
	auth := router.Group("/user")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	{
		auth.GET("/:id", getUserProfile)
		auth.PATCH("/:id", updateUserProfile)
		auth.DELETE("/:id", deleteUserProfile)
	}

	router.Run("localhost:" + port)
}
