package main

import (
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Registers a new user
func registerUserProfile(c *gin.Context) {
	var payload UserProfile

	// Reads JSON Body
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Username == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
		return
	}

	// Makes new database entry
	err = DB.Create(&UserProfile{
		Username: payload.Username,
		Password: payload.Password,
	}).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered Successfully"})
}

// Gets the profile of a user
func getUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user UserProfile
	queryResult := DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Updates data for an already existing user
func updateUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user UserProfile
	queryResult := DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Reads JSON body
	var payload UserProfile
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Username == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
		return
	}

	// Updates Profile
	user.Username = payload.Username
	user.Password = payload.Password

	err = DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})
}

// Deletes a user
func deleteUserProfile(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing User ID"})
		return
	}

	// Query user from database
	var user UserProfile
	queryResult := DB.First(&user, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Deletes user
	var err error = DB.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}

// Creates a post
func createPost(c *gin.Context) {

	// Gets id set in JWT token
	claims := jwt.ExtractClaims(c)
	idRaw, ok := claims["id"]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 1"})
		return
	}

	// Convert to uint (because your DB expects uint)
	var userId uint
	switch v := idRaw.(type) {
	case float64:
		userId = uint(v)

	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 2"})
			return
		}

		userId = uint(parsed)

	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 3"})
		return
	}

	var payload Post
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post cannot be empty"})
		return
	}

	// Creates new post
	post := Post{
		Content:       payload.Content,
		UserProfileID: userId,
	}

	err = DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// Gets a post
func getPost(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Query user from database
	var p Post
	queryResult := DB.First(&p, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, p)
}

// Deletes Post
func deletePost(c *gin.Context) {
	reqId := c.Param("id")
	if reqId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing Post ID"})
		return
	}

	// Query user from database
	var p Post
	queryResult := DB.First(&p, "id = ?", reqId)
	if queryResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Deletes user
	var err error = DB.Delete(&p).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
