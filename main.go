/*
Next Steps:
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

	initDB(os.Getenv("DB_URI")) // Initialises Database

	// Starts Gin router
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
	auth := router.Group("/")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	{
		// User routes
		auth.GET("/users/:id", getUserProfile)
		auth.PATCH("/users/:id", updateUserProfile)
		auth.DELETE("/users/:id", deleteUserProfile)

		// Post routes
		auth.POST("/posts", createPost)
		auth.GET("/posts/:pid", getPost)
		auth.DELETE("/posts/:pid", deletePost)

		// Comment routes
		auth.POST("/posts/:pid/comments", createComment)
		auth.GET("/posts/:pid/comments/:cid", getComment)
		auth.GET("/posts/:pid/comments/all", getCommentsAll)
		auth.DELETE("/posts/:pid/comments/:cid", deleteComment)
	}

	router.Run("localhost:" + os.Getenv("PORT"))
}
