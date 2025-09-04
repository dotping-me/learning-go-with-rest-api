package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserProfile(c *gin.Context) {
	reqId := c.Query("id")
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

func registerUserProfile(c *gin.Context) {
	var payload UserProfile

	// Reads JSON Body
	var err error = c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if payload.Email == "" || payload.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
		return
	}

	// Makes new database entry
	DB.Create(&UserProfile{
		Email:    payload.Email,
		Username: payload.Username,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Registered Successfully"})
}

func updateUserProfile(c *gin.Context) {
	reqId := c.Query("id")
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

	if payload.Email == "" || payload.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON Body missing some fields"})
		return
	}

	// Updates Profile
	user.Email = payload.Email
	user.Username = payload.Username

	err = DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})
}
