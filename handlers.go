package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, user) // Sends JSON response
}

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
	DB.Create(&UserProfile{
		Username: payload.Username,
		Password: payload.Password,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Registered Successfully"})
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
