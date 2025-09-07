/*

API Routes are registered here and mapped to respective
handler functions!

*/

package api

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {

	// Unprotected routes
	router.POST("api/v1/login", jwtMiddleware.LoginHandler) // JWT Token Interaction
	router.POST("api/v1/signup", registerUserProfile)

	// Protected routes
	auth := router.Group("/api/v1")
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
}

func RegisterWebRoutes(router *gin.Engine, jwtMiddleware *jwt.GinJWTMiddleware) {
	router.GET("/", HomePage)
	router.GET("/login", LoginPage)
}
